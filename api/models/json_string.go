// models/json_string.go
package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONString map[string]interface{}

func (j JSONString) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

func (j *JSONString) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}
	var v map[string]interface{}
	err := json.Unmarshal(bytes, &v)
	if err != nil {
		return err
	}
	*j = v
	return nil
}
