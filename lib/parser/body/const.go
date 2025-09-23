package bodyparser

const (
	MIMEJSON              = "application/json"
	MIMEHTML              = "text/html"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
	MIMEPROTOBUF          = "application/x-protobuf"
	MIMEMSGPACK           = "application/x-msgpack"
	MIMEMSGPACK2          = "application/msgpack"
	MIMEYAML              = "application/x-yaml"
	MIMETOML              = "application/toml"
)

const (
	JSON = "json"
	FORM = "form"
)

const (
	Required        = "required"
	RequiredWithout = "required_without"
	RequiredWith    = "required_with"
	RequiredIf      = "required_if"
	Max             = "max"
	Min             = "min"
	EqField         = "eqfield"
	Email           = "email"
	Len             = "len"
)
