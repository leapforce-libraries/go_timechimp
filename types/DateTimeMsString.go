package timechimp

import (
	"encoding/json"
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
)

const (
	dateTimeMsFormat string = "2006-01-02T15:04:05.999"
)

type DateTimeMsString time.Time

func (d *DateTimeMsString) UnmarshalJSON(b []byte) error {
	var returnError = func() error {
		errortools.CaptureError(fmt.Sprintf("Cannot parse '%s' to DateTimeMsString", string(b)))
		return nil
	}

	var s string

	err := json.Unmarshal(b, &s)
	if err != nil {
		return returnError()
	}

	if len(s) > len(dateTimeMsFormat) {
		s = s[:len(dateTimeMsFormat)]
	}

	if s == "" || s == "0000-00-00 00:00:00.000" {
		d = nil
		return nil
	}

	_t, err := time.Parse(dateTimeMsFormat, s)
	if err != nil {
		return returnError()
	}

	*d = DateTimeMsString(_t)
	return nil
}

func (d *DateTimeMsString) ValuePtr() *time.Time {
	if d == nil {
		return nil
	}

	_d := time.Time(*d)
	return &_d
}

func (d DateTimeMsString) Value() time.Time {
	return time.Time(d)
}

func (d DateTimeMsString) String() string {
	return time.Time(d).Format(dateTimeMsFormat)
}
