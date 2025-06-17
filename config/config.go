package config

const (
	DEFAULT_TIME_ZONE = "DEFAULT_TIME_ZONE"
	DEFAULT_PASSWORD  = "DEFAULT_PASSWORD"
	CLIENT_SECRET     = "CLIENT_SECRET"
)

var basicConfig = map[string]string{
	ENV:               ENV_DEVELOPMENT,
	ENV_DIR:           ENV_ROOT_DIR,
	DEFAULT_TIME_ZONE: "Asia/Jakarta",
	DEFAULT_PASSWORD:  "password",
	CLIENT_SECRET:     "local_secret",
}

var base = mergeConfig(
	authConfig,
	basicConfig,
	postgresConfig,
	httpConfig,
)

var baseInterface = mergeConfigInterface(
	appInterfaceConfig,
	httpInterfaceConfig,
)
