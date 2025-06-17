package config

const (
	JWT_PRIVATE_KEY = "JWT_PRIVATE_KEY"
	JWT_PUBLIC_KEY  = "JWT_PUBLIC_KEY"
)

var authConfig = map[string]string{
	JWT_PRIVATE_KEY: "",
	JWT_PUBLIC_KEY:  "",
}
