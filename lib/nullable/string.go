package nullable

import (
	"encoding/json"
)

type String struct {
	Value *string `json:"value"`
	Valid bool    `json:"valid"`
}

func (i *String) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		i.Value = nil
		i.Valid = true
		return nil
	}

	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	i.Value = &value
	i.Valid = true
	return nil
}
