// Package config hanldes of config to run the app
package config

const (
	DefaultTimeZone = "DefaultTimeZone"
	DefaultPassword = "DefaultPassword"
	ClientSecret    = "ClientSecret"
)

var basicConfig = map[string]string{
	Env:             EnvDevelopment,
	EnvDir:          EnvRootDir,
	DefaultTimeZone: "Asia/Jakarta",
	DefaultPassword: "password",
	ClientSecret:    "local_secret",
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
