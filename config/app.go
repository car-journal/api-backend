package config

import "time"

var appInterfaceConfig = map[string]interface{}{}

const (
	SERVICE_NAME      = "SERVICE_NAME"
	ENV               = "ENV"
	ENV_DIR           = "ENV_DIR"
	ENCRYPTION_SECRET = "ENCRYPTION_SECRET"
	LOCALE            = "LOCALE"
)

var DEFAULT_START_DATE = time.Now().AddDate(0, -1, 0)
var DEFAULT_LIMIT = 20
