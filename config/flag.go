package config

import "github.com/car-journal/api-backend/lib/flag"

const (
	FeatureLog = "FEATURE_LOG"
)

var flagConfig = map[string]string{
	FeatureLog: flag.FeatureOn,
}

func GetFlagConfig() map[string]string {
	return flagConfig
}
