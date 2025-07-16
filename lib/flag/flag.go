// Package flag handles flagging system
package flag

import (
	"os"
)

type FlagInterface interface {
	Flag(key string) bool
}

type flag struct {
	featureFlagCfg map[string]string
}

func NewFlag(
	featureFlagCfg map[string]string,
) FlagInterface {
	return &flag{
		featureFlagCfg: featureFlagCfg,
	}
}

func (f flag) Flag(key string) bool {
	return f.get(key) == FeatureOn
}

func (f flag) get(key string) string {
	r := os.Getenv(key)
	if r != "" {
		return r
	}
	if configValue, ok := f.featureFlagCfg[key]; ok {
		return configValue
	}
	return ""
}
