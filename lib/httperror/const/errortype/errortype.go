// Package errortype handles constants for error types
package errortype

type ErrorType string

func (e ErrorType) ToString() string {
	return string(e)
}

const (
	Forbidden           ErrorType = "forbidden"
	InternalServer      ErrorType = "internal_server"
	InvalidInput        ErrorType = "invalid_input"
	RecordNotFound      ErrorType = "record_not_found"
	InputRecordNotFound ErrorType = "input_record_not_found"
	NotAcceptable       ErrorType = "not_acceptable"
	AlreadyRegistered   ErrorType = "already_registered"
	Unauthenticated     ErrorType = "unauthentication"
	Unauthorized        ErrorType = "unauthorized"
	PaymentRequired     ErrorType = "payment_required"
	MissingHeader       ErrorType = "missing_header"
	InvalidHeader       ErrorType = "invalid_header"
)
