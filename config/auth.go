package config

const (
	JwtPrivateKey = "JwtPrivateKey"
	JwtPublicKey  = "JwtPublicKey"
)

var authConfig = map[string]string{
	JwtPrivateKey: "",
	JwtPublicKey:  "",
}
