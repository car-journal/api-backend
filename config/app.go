package config

import "time"

var appInterfaceConfig = map[string]interface{}{}

const (
	ServiceName      = "ServiceName"
	Env              = "ENV"
	EnvDir           = "ENV_DIR"
	EncryptionSecret = "ENCRYPTION_SECRET"
	Locale           = "Locale"
)

var DefaultStartTime = time.Now().AddDate(0, -1, 0)
var DefaultLimit = 20
