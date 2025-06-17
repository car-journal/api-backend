package nullable

import (
	"encoding/json"
)

type Int64 struct {
	Value *int64 `json:"value"`
	Valid bool   `json:"valid"`
}

func (i *Int64) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		i.Value = nil
		i.Valid = true
		return nil
	}

	var value int64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	i.Value = &value
	i.Valid = true
	return nil
}
