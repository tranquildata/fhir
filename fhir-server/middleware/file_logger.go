package middleware

import (
	"strconv"
	"bytes"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/DataDog/zstd"
	"github.com/pkg/errors"
	"golang.org/x/sys/unix"

	"github.com/gin-gonic/gin"
)

type responseTeeWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseTeeWriter) Write(b []byte) (int, error) {
	return w.body.Write(b)
}

/*
	Writes raw requests and responses to files for archiving.
	Compression using Zstandard (which is famous for high compression speed)

	Sets several xattrs: http_status, latency_ms, requestor_details, mutex_name
*/
func FileLoggerMiddleware(outputDirectory string, dumpHttpGET bool) gin.HandlerFunc {

	// writing to temp directory and then moving into outputDirectory
	// so that monitoring tools don't see a partially-written file
	tempDirectory := path.Join(outputDirectory, "temp")

	os.MkdirAll(tempDirectory, 0777)

	return func(c *gin.Context) {

		if !dumpHttpGET && c.Request.Method == "GET" {
			c.Next()
			return
		}

		mutexName := c.GetHeader("X-Mutex-Name")
		requestorDetails := c.GetHeader("X-Requestor-Details")
		db := c.GetHeader("Db")
		if db != "" {
			// assume re-entrant call (via HandleContext() in routing.go)
			c.Next()
			return
		}

		// read request
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// form filenames
		hasher := sha1.New()
		hasher.Write(bodyBytes)
		hashBytes := hasher.Sum(nil)
		currentTime := time.Now()
		timestamp := currentTime.Format("2006-01-02-15-04-05.000000")
		filename := fmt.Sprintf("%s-%x", timestamp, hashBytes)
		requestFilename := filename+".req.zst"
		responseFilename := filename+".res.zst"

		// write request
		requestWritten := make(chan error)
		go func() {

			ferr := writeCompressedData(tempDirectory, requestFilename, bodyBytes)
			if ferr != nil {
				requestWritten <- ferr
				return
			}

			ferr = moveToDir(requestFilename, tempDirectory, outputDirectory)
			if ferr != nil {
				requestWritten <- ferr
				return
			}

			requestWritten <- nil
		}()

		// hook the response
		tee := &responseTeeWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = tee

		// do the work
		started := time.Now()
		c.Next()
		latencyMsecs := time.Since(started).Nanoseconds() / 1e6

		// write the response
		statusCode := c.Writer.Status()


		err := writeCompressedData(tempDirectory, responseFilename, tee.body.Bytes())
		if err != nil {
			c.AbortWithError(500, errors.Wrap(err, "FileLoggerMiddleware: writeCompressedData failed"))
			return
		}

		err = setXattrInt(tempDirectory, responseFilename, "user.http_status", int64(statusCode), c)
		if err != nil {
			return
		}
		err = setXattrInt(tempDirectory, responseFilename, "user.latency_ms", latencyMsecs, c)
		if err != nil {
			return
		}
		err = setXattr(tempDirectory, responseFilename, "user.mutex_name", mutexName, c)
		if err != nil {
			return
		}
		err = setXattr(tempDirectory, responseFilename, "user.requestor_details", requestorDetails, c)
		if err != nil {
			return
		}

		err = moveToDir(responseFilename, tempDirectory, outputDirectory)
		if err != nil {
			c.AbortWithError(500, errors.Wrapf(err, "FileLoggerMiddleware: os.Rename for response failed: %s", responseFilename))
			return
		}

		// wait for request to be written
		err = <-requestWritten
		if err != nil {
			c.AbortWithError(500, errors.Wrap(err, "FileLoggerMiddleware: failed to write request"))
			return
		}

		// write response
		c.Header("X-GoFHIR-Request-ID", filename)
		tee.ResponseWriter.Write(tee.body.Bytes())
	}
}

func setXattr(tempDirectory string, filename string, xattrName string, xattrValue string, c *gin.Context) error {
	if xattrValue == "" {
		return nil
	}
	err := unix.Setxattr(path.Join(tempDirectory, filename), xattrName, []byte(xattrValue), 0)
	if err != nil {
		c.AbortWithError(500, errors.Wrapf(err, "FileLoggerMiddleware: setxattr (%s) (%d bytes) failed on %s", xattrName, len(xattrValue), filename))
	}
	return err
}
func setXattrInt(tempDirectory string, filename string, xattrName string, xattrValue int64, c *gin.Context) error {
	return setXattr(tempDirectory, filename, xattrName, strconv.FormatInt(xattrValue, 10), c)
}
func moveToDir(filename string, fromDir string, toDir string) error {
	return os.Rename(path.Join(fromDir, filename), path.Join(toDir, filename))
}
func writeCompressedData(dirName string, filename string, data []byte) error {
	f, ferr := os.Create(path.Join(dirName, filename))
	if ferr != nil {
		return errors.Wrap(ferr, "os.Create failed")
	}
	defer f.Close()

	compressor := zstd.NewWriterLevel(f, 5)
	_, ferr = compressor.Write(data)
	if ferr != nil {
		return errors.Wrap(ferr, "compressor.Write failed")
	}
	ferr = compressor.Close()
	if ferr != nil {
		return errors.Wrap(ferr, "compressor.Close failed")
	}

	return nil
}