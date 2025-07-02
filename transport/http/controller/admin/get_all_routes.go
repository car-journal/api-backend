package admincontroller

import (
	"net/http"

	"github.com/car-journal/api-backend/lib/parser"
	"github.com/gorilla/mux"
)

func GetAllRoutes(route *mux.Router) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		type responseStruct struct {
			Method  string `json:"method"`
			Path    string `json:"path"`
			Handler string `json:"handler"`
		}
		var response []responseStruct
		if err := route.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			path, err := route.GetPathTemplate()
			if err != nil {
				return err
			}

			methods, err := route.GetMethods()
			var method string
			if err != nil && err.Error() != "mux: route doesn't have methods" {
				return err
			}
			if len(methods) > 0 {
				method = methods[0]
			}

			name := route.GetName()

			response = append(response, responseStruct{
				Path:    path,
				Method:  method,
				Handler: name,
			})
			return nil
		}); err != nil {
			parser.JSON(writer, nil, err)
			return
		}

		parser.Json(writer, response, nil)
	}
}
