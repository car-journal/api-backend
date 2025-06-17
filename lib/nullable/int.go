package nullable

import (
	"encoding/json"
)

type Int struct {
	Value *int `json:"value"`
	Valid bool `json:"valid"`
}

func (i *Int) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		i.Value = nil
		i.Valid = true
		return nil
	}

	var value int
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	i.Value = &value
	i.Valid = true
	return nil
}
