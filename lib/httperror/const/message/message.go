// Package message handles constants for errors messages
package message

type Message string

// TODO: make appropriate messages
const (
	Forbidden           Message = "forbidden"
	InternalServer      Message = "something went wrong in our end; please try again later"
	InvalidInput        Message = "please re-check your input; make sure the validations are met"
	RecordNotFound      Message = "enter another param or recheck the param you've sent"
	InputRecordNotFound Message = "enter another input value or recheck the value you've sent"
	NotAcceptable       Message = "not_acceptable"
	AlreadyRegistered   Message = "this credential already exists; try another one"
	Unauthenticated     Message = "unauthentication"
	Unauthorized        Message = "your credentials is unauthorized; please re-check your credential"
	PaymentRequired     Message = "payment_required"
	MissingHeader       Message = "missing_header"
	InvalidHeader       Message = "invalid_header"
	Default             Message = "check everything"
)

func (m Message) ToString() string {
	return string(m)
}
