// Package httperror handles error interface
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

	errorType := errortype.InternalServer
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errorType = errortype.RecordNotFound
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
	res.Message = message.Default
	switch errorType {
	case errortype.Forbidden:
		res.Code = http.StatusForbidden
		res.Message = message.Forbidden
	case errortype.InternalServer:
		res.Code = http.StatusInternalServerError
		res.Message = message.InternalServer
	case errortype.InvalidInput:
		res.Message = message.InvalidInput
	case errortype.RecordNotFound:
		res.Code = http.StatusNotFound
		res.Message = message.RecordNotFound
	case errortype.InputRecordNotFound:
		res.Message = message.InputRecordNotFound
	case errortype.NotAcceptable:
		res.Code = http.StatusNotAcceptable
		res.Message = message.NotAcceptable
	case errortype.AlreadyRegistered:
		res.Code = http.StatusConflict
		res.Message = message.AlreadyRegistered
	case errortype.Unauthenticated:
		res.Code = http.StatusUnauthorized
		res.Message = message.Unauthenticated
	case errortype.Unauthorized:
		res.Code = http.StatusUnauthorized
		res.Message = message.Unauthorized
	case errortype.PaymentRequired:
		res.Code = http.StatusPaymentRequired
		res.Message = message.PaymentRequired
	case errortype.MissingHeader:
		res.Message = message.MissingHeader
	case errortype.InvalidHeader:
		res.Message = message.InvalidHeader
	default:
		res.Message = message.Default
	}

	return res
}
