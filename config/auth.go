package config

const (
	JwtPrivateKey = "JWT_PRIVATE_KEY"
	JwtPublicKey  = "JWT_PUBLIC_KEY"
)

var authConfig = map[string]string{
	JwtPrivateKey: "",
	JwtPublicKey:  "",
}
