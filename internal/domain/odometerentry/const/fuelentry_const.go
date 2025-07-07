// Package fuelentryconst handles const for fuel entry domain
package fuelentryconst

import "strings"

const (
	MALE   = "male"
	FEMALE = "female"
)

func GetGenderResponse(genderDB bool) string {
	if genderDB {
		return MALE
	}

	return FEMALE
}

func GetGenderDB(genderRequest *string) *bool {
	if genderRequest == nil {
		return nil
	}

	res := strings.ToLower(*genderRequest) == MALE
	return &res
}
