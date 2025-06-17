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
	REQUIRED         = "required"
	REQUIRED_WITHOUT = "required_without"
	REQUIRED_WITH    = "required_with"
	REQUIRED_IF      = "required_if"
	MAX              = "max"
	MIN              = "min"
	EQ_FIELD         = "eqfield"
	EMAIL            = "email"
	LEN              = "len"
)
