package httperror

import (
	"errors"
	"net/http"

	"github.com/car-journal/api-backend/lib/httperror/const/errortype"
	"github.com/car-journal/api-backend/lib/httperror/const/message"
	"gorm.io/gorm"
)

type Interface interface {
	Error() string
	GetMessage() message.Message
	GetCode() int
	GetType() string
	GetRaw() interface{}
}

type Base struct {
	Type       errortype.ErrorType `json:"type"`
	Code       int                 `json:"-"`
	StatusCode string              `json:"code"`
	Raw        interface{}         `json:"raw"`
	Message    message.Message     `json:"message"`
}

func (e *Base) Error() string {
	if e.Message == "" {
		if result, ok := e.Raw.(error); ok {
			return result.Error()
		}
	}
	return e.Message.ToString()
}

func (e *Base) GetMessage() message.Message {
	return e.Message
}

func (e *Base) GetCode() int {
	return e.Code
}

func (e *Base) GetType() string {
	return e.Type.ToString()
}

func (e *Base) GetRaw() interface{} {
	if result, ok := e.Raw.(error); ok {
		return result.Error()
	}

	return e.Raw
}

func New(errorType errortype.ErrorType, raw interface{}) Interface {
	err, ok := raw.(error)
	if ok && err == nil {
		return nil
	}

	codeMessage := IssueCodeMessage(errorType)

	return &Base{
		Type:       errorType,
		Code:       codeMessage.Code,
		StatusCode: http.StatusText(codeMessage.Code),
		Message:    codeMessage.Message,
		Raw:        raw,
	}
}

func GetInstance(err error) Interface {
	if result, ok := err.(*Base); ok {
		return result
	}

	errorType := errortype.INTERNAL_SERVER
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errorType = errortype.RECORD_NOT_FOUND
	}
	codeMessage := IssueCodeMessage(errorType)

	return &Base{
		Type:       errorType,
		Code:       codeMessage.Code,
		StatusCode: http.StatusText(codeMessage.Code),
		Message:    codeMessage.Message,
		Raw:        err,
	}
}

type CodeMessage struct {
	Code            int `json:"code"`
	message.Message `json:"message"`
}

func IssueCodeMessage(errorType errortype.ErrorType) (res CodeMessage) {
	res.Code = http.StatusBadRequest
	res.Message = message.DEFAULT
	switch errorType {
	case errortype.FORBIDDEN:
		res.Code = http.StatusForbidden
		res.Message = message.FORBIDDEN
	case errortype.INTERNAL_SERVER:
		res.Code = http.StatusInternalServerError
		res.Message = message.INTERNAL_SERVER
	case errortype.INVALID_INPUT:
		res.Message = message.INVALID_INPUT
	case errortype.RECORD_NOT_FOUND:
		res.Code = http.StatusNotFound
		res.Message = message.RECORD_NOT_FOUND
	case errortype.INPUT_RECORD_NOT_FOUND:
		res.Message = message.INPUT_RECORD_NOT_FOUND
	case errortype.NOT_ACCEPTABLE:
		res.Code = http.StatusNotAcceptable
		res.Message = message.NOT_ACCEPTABLE
	case errortype.ALREADY_REGISTERED:
		res.Code = http.StatusConflict
		res.Message = message.ALREADY_REGISTERED
	case errortype.UNAUTHENTICATED:
		res.Code = http.StatusUnauthorized
		res.Message = message.UNAUTHENTICATED
	case errortype.UNAUTHORIZED:
		res.Code = http.StatusUnauthorized
		res.Message = message.UNAUTHORIZED
	case errortype.PAYMENT_REQUIRED:
		res.Code = http.StatusPaymentRequired
		res.Message = message.PAYMENT_REQUIRED
	case errortype.MISSING_HEADER:
		res.Message = message.MISSING_HEADER
	case errortype.INVALID_HEADER:
		res.Message = message.INVALID_HEADER
	default:
		res.Message = message.DEFAULT
	}

	return res
}
