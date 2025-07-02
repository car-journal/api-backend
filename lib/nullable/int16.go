package nullable

import (
	"encoding/json"
)

type Int16 struct {
	Value *int16 `json:"value"`
	Valid bool   `json:"valid"`
}

func (i *Int16) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		i.Value = nil
		i.Valid = true
		return nil
	}

	var value int16
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	i.Value = &value
	i.Valid = true
	return nil
}
