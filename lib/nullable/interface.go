package nullable

import (
	"encoding/json"
)

type Interface struct {
	Value []interface{} `json:"value"`
	Valid bool          `json:"valid"`
}

func (ife *Interface) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		ife.Value = nil
		ife.Valid = true
		return nil
	}

	var value []interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	ife.Value = value
	ife.Valid = true
	return nil
}
