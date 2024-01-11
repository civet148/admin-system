package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/civet148/log"
	"reflect"
)

type Storage struct {
	StorageID string   `json:"storage_id" db:"storage_id" bson:"storage_id"` //storage id
	Unsealed  bool     `json:"unsealed" db:"unsealed" bson:"unsealed"`       //unsealed
	Sealed    bool     `json:"sealed" db:"sealed" bson:"sealed"`             //sealed
	Cache     bool     `json:"cache" db:"cache" bson:"cache"`                //cache
	Local     string   `json:"local" db:"local" bson:"local"`                //local storage path
	URLs      []string `json:"urls" db:"urls" bson:"urls"`                   //remote access urls
}
type StorageSlice []*Storage

func (s StorageSlice) GoString() string {
	return s.String()
}

func (s StorageSlice) String() string {
	data, _ := json.MarshalIndent(s, "", "\t")
	return string(data)
}

func (s StorageSlice) Len() int {
	return len(s)
}

func (s StorageSlice) ForEach(cb func(*Storage) error) {
	for _, v := range s {
		if err := cb(v); err != nil {
			break
		}
	}
}

// Scan implements the sql.Scanner interface for database deserialization.
func (s StorageSlice) Scan(src interface{}) error {
	var data []byte
	var err error

	switch src.(type) {
	case []byte:
		data = src.([]byte)
	case string:
		data = []byte(src.(string))
	case *string:
		data = []byte(*src.(*string))
	default:
		err = fmt.Errorf("can not handle with unknown type [%v], just only []byte or string supported", reflect.TypeOf(src).Kind())
		log.Errorf(err.Error())
		return err
	}
	err = json.Unmarshal(data, s)
	return err
}

// Value implements the driver.Valuer interface for database serialization.
func (s StorageSlice) Value() (value driver.Value, err error) {
	var data []byte
	data, err = json.MarshalIndent(s, "", "\t")
	return string(data), nil
}
