package models

import (
	"fmt"
	"encoding/json"
	"time"
)

type Precision string

const (
	Date      = "date"
	YearMonth = "year-month"
	Year = "year"
	Timestamp = "timestamp"
	Time      = "time"
)

type FHIRDateTime struct {
	Time      time.Time
	Precision Precision
}

func (f *FHIRDateTime) UnmarshalJSON(data []byte) (err error) {
	strData := string(data)
	if len(data) <= 12 {
		f.Precision = Precision("date")
		f.Time, err = time.ParseInLocation("\"2006-01-02\"", strData, time.Local)
		if err != nil {
			f.Precision = Precision("year-month")
			f.Time, err = time.ParseInLocation("\"2006-01\"", strData, time.Local)
		}
		if err != nil {
			f.Precision = Precision("year")
			f.Time, err = time.ParseInLocation("\"2006\"", strData, time.Local)
		}
		if err != nil {
			// TODO: should move time into a separate type
			f.Precision = Precision("time")
			f.Time, err = time.ParseInLocation("\"15:04:05\"", strData, time.Local)
		}
		if err != nil {
			err = fmt.Errorf("unable to parse DateTime: %s", strData)
			f.Precision = ""
		}

	} else {
		f.Precision = Precision("timestamp")
		f.Time = time.Time{}
		f.Time.UnmarshalJSON(data)
	}
	return err
}

func (f FHIRDateTime) MarshalJSON() ([]byte, error) {
	if f.Precision == Timestamp {
		return json.Marshal(f.Time.Format(time.RFC3339))
	} else if f.Precision == YearMonth {
		return json.Marshal(f.Time.Format("2006-01"))
	} else if f.Precision == Year {
		return json.Marshal(f.Time.Format("2006"))
	} else if f.Precision == Time {
		return json.Marshal(f.Time.Format("15:04:05"))
	} else {
		return json.Marshal(f.Time.Format("2006-01-02"))
	}
}
