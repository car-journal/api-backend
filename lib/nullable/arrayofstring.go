package nullable

import "encoding/json"

type ArrayOfString struct {
	Value []string `json:"value"`
	Valid bool     `json:"valid"`
}

func (aos *ArrayOfString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		aos.Value = nil
		aos.Valid = true
		return nil
	}

	var value []string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	aos.Value = value
	aos.Valid = true
	return nil
}
