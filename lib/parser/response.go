package parser

import (
	"encoding/json"
	"net/http"

	"github.com/car-journal/api-backend/lib/httperror"
	"github.com/car-journal/api-backend/lib/httperror/const/message"
)

type response struct {
	Success bool           `json:"success"`
	Error   *responseError `json:"error,omitempty"`
}

type responseError struct {
	Type    string          `json:"type,omitempty"`
	Message message.Message `json:"message,omitempty"`
	Raw     interface{}     `json:"raw"`
}

func JSON(w http.ResponseWriter, data interface{}, err error) {
	w.Header().Set("Content-Type", "application/json")
	var response response
	code := http.StatusOK
	if nil != err {
		errInstance := httperror.GetInstance(err)
		code = errInstance.GetCode()
		response.Success = false
		response.Error = &responseError{
			Type:    errInstance.GetType(),
			Message: errInstance.GetMessage(),
			Raw:     errInstance.GetRaw(),
		}
		w.WriteHeader(code)
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(code)
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
		return
	}

	response.Success = true
	_ = json.NewEncoder(w).Encode(response)
}
