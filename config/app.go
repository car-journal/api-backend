package config

import "time"

var appInterfaceConfig = map[string]interface{}{}

const (
	ServiceName      = "ServiceName"
	Env              = "Env"
	EnvDir           = "EnvDir"
	EncryptionSecret = "EncryptionSecret"
	Locale           = "Locale"
)

var DefaultStartTime = time.Now().AddDate(0, -1, 0)
var DefaultLimit = 20
