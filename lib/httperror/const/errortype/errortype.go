package errortype

type ErrorType string

func (e ErrorType) ToString() string {
	return string(e)
}

const (
	FORBIDDEN              ErrorType = "forbidden"
	INTERNAL_SERVER        ErrorType = "internal_server"
	INVALID_INPUT          ErrorType = "invalid_input"
	RECORD_NOT_FOUND       ErrorType = "record_not_found"
	INPUT_RECORD_NOT_FOUND ErrorType = "input_record_not_found"
	NOT_ACCEPTABLE         ErrorType = "not_acceptable"
	ALREADY_REGISTERED     ErrorType = "already_registered"
	UNAUTHENTICATED        ErrorType = "unauthentication"
	UNAUTHORIZED           ErrorType = "unauthorized"
	PAYMENT_REQUIRED       ErrorType = "payment_required"
	MISSING_HEADER         ErrorType = "missing_header"
	INVALID_HEADER         ErrorType = "invalid_header"
)
