package types

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strings"
	"time"
)

const (
	TIME_FORMAT_DATE = "2006-01-02"
)

type Date string

func (d Date) String() string {
	return string(d)
}

// MarshalJSON implements the json.Marshaler interface.
func (d Date) MarshalJSON() ([]byte, error) {
	strDateWithQuote := fmt.Sprintf("\"%s\"", d)
	return []byte(strDateWithQuote), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *Date) UnmarshalJSON(decimalBytes []byte) error {
	var t time.Time
	var dateLenIncludeQuotes = 12 // date string include quotes, e.g "2021-05-08" as []byte length is 12 not 10
	if len(decimalBytes) > dateLenIncludeQuotes {
		//maybe is a timestamp like "2021-05-08T13:03:22Z"
		if err := t.UnmarshalJSON(decimalBytes); err != nil {
			return err
		}
		*d = Date(t.Format(TIME_FORMAT_DATE))
	} else {
		strDate := strings.Replace(string(decimalBytes), "\"", "", -1) //trim quotes
		*d = Date(strDate)
	}
	return nil
}

// Scan implements the sql.Scanner interface for database deserialization.
func (d *Date) Scan(src interface{}) (err error) {
	var value []byte
	switch src.(type) {
	case int, int32, int64, uint, uint32, uint64, string:
		value = []byte(fmt.Sprintf("%v", src))
	case []byte:
		value = src.([]byte)
	default:
		err = fmt.Errorf("unknown type [%v] to scan", reflect.TypeOf(src).String())
		fmt.Printf("%s\n", err)
		return
	}
	if err = d.UnmarshalJSON(value); err != nil {
		fmt.Printf("unmarshal [%s] error [%s]\n", value, err)
		return
	}
	return
}

// Value implements the driver.Valuer interface for database serialization.
func (d Date) Value() (driver.Value, error) {
	return d.String(), nil
}
