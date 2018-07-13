# Old build using go get:
# COPY . /go/src/github.com/eug48/fhir/
# RUN ls -l /go/src/github.com/eug48/fhir/
# WORKDIR /
# RUN go get -d -v github.com/eug48/fhir/fhir-server
# WORKDIR /go/src/github.com/eug48/fhir/fhir-server
# RUN CGO_ENABLED=0 GOOS=linux go build

# Build using Dep
FROM golang:1.10.2 as builder
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 && chmod +x /usr/local/bin/dep

# Get dependencies first for docker caching
WORKDIR /go/src/github.com/eug48/fhir/
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only -v

# Copy source
COPY . ./

# Build
WORKDIR /go/src/github.com/eug48/fhir/fhir-server
RUN CGO_ENABLED=0 GOOS=linux go build


# Copy to light-weight runtime image
FROM alpine:3.7
RUN apk add --no-cache ca-certificates tini
COPY --from=builder /go/src/github.com/eug48/fhir/fhir-server/fhir-server /
COPY --from=builder /go/src/github.com/eug48/fhir/fhir-server/config/ /config
COPY --from=builder /go/src/github.com/eug48/fhir/conformance/ /conformance

ENV MONGO_HOSTPORT fhir-mongo:27017
CMD ["sh", "-c", "/fhir-server -port 3001 -enableXML -databaseName fhir -mongodbHostPort $MONGO_HOSTPORT"]