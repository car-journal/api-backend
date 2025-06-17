package message

type Message string

// TODO: make appropriate messages
const (
	FORBIDDEN              Message = "forbidden"
	INTERNAL_SERVER        Message = "something went wrong in our end; please try again later"
	INVALID_INPUT          Message = "please re-check your input; make sure the validations are met"
	RECORD_NOT_FOUND       Message = "enter another param or recheck the param you've sent"
	INPUT_RECORD_NOT_FOUND Message = "enter another input value or recheck the value you've sent"
	NOT_ACCEPTABLE         Message = "not_acceptable"
	ALREADY_REGISTERED     Message = "this credential already exists; try another one"
	UNAUTHENTICATED        Message = "unauthentication"
	UNAUTHORIZED           Message = "your credentials is unauthorized; please re-check your credential"
	PAYMENT_REQUIRED       Message = "payment_required"
	MISSING_HEADER         Message = "missing_header"
	INVALID_HEADER         Message = "invalid_header"
	DEFAULT                Message = "check everything"
)

func (m Message) ToString() string {
	return string(m)
}
