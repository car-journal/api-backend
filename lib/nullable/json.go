package nullable

import (
	"encoding/json"

	"gorm.io/datatypes"
)

type JSON struct {
	Value *datatypes.JSON `json:"value"`
	Valid bool            `json:"valid"`
}

func (j *JSON) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		j.Value = nil
		j.Valid = true
		return nil
	}

	var value *datatypes.JSON
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	j.Value = value
	j.Valid = true
	return nil
}
