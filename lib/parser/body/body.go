// Package bodyparser handles request body parsing
package bodyparser

import (
	"encoding/json"
	"net/http"

	carjournalstrings "github.com/car-journal/api-backend/lib/strings"
	"github.com/go-playground/form/v4"
)

func ParseForm(r *http.Request, i interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	return form.NewDecoder().Decode(i, r.Form)
}

func ParseJSON(r *http.Request, i interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(i)
}

func Parse(r *http.Request, i interface{}) (err error) {
	b := Default(r.Method, r.Header.Get("Content-type"))
	switch b {
	case JSON:
		err = ParseJSON(r, i)
	default:
		err = ParseForm(r, i)
	}
	if nil != err {
		return err
	}

	return Validate(i, b)
}

func Default(method, contentType string) string {
	if method == http.MethodGet {
		return FORM
	}

	array := carjournalstrings.Split(contentType, ";")
	if len(array) <= 0 {
		return FORM
	}

	switch array[0] {
	case MIMEJSON:
		return JSON
	// case MIMEXML, MIMEXML2:
	// 	return XML
	// case MIMEPROTOBUF:
	// 	return ProtoBuf
	// case MIMEMSGPACK, MIMEMSGPACK2:
	// 	return MsgPack
	// case MIMEYAML:
	// 	return YAML
	// case MIMETOML:
	// 	return TOML
	// case MIMEMultipartPOSTForm:
	// 	return FormMultipart
	default: // case MIMEPOSTForm:
		return FORM
	}
}
