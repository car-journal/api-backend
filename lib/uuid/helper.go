// Package uuid helps to have a single file to handle all uuid values
package uuid

import (
	gofrsUUID "github.com/gofrs/uuid"
)

func MustStringToUUID(s string) UUID {
	res, err := StringToUUID(s)
	return UUID{
		UUID: gofrsUUID.Must(res.UUID, err),
	}
}

func StringToUUID(s string) (UUID, error) {
	res, err := gofrsUUID.FromString(s)
	return UUID{
		UUID: res,
	}, err
}

func IsValidUUID(s string) bool {
	_, err := gofrsUUID.FromString(s)
	return err == nil
}
