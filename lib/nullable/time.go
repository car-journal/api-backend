package nullable

import (
	"encoding/json"
	"time"
)

type Time struct {
	Value *time.Time `json:"value"`
	Valid bool       `json:"valid"`
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		t.Value = nil
		t.Valid = true
		return nil
	}

	var value time.Time
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	t.Value = &value
	t.Valid = true
	return nil
}
