package nullable

import (
	"encoding/json"
)

type Float64 struct {
	Value *int64 `json:"value"`
	Valid bool   `json:"valid"`
}

func (f *Float64) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		f.Value = nil
		f.Valid = true
		return nil
	}

	var value int64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	f.Value = &value
	f.Valid = true
	return nil
}
