package http

import (
	"encoding/json"
	"net/http"

	"github.com/asinha24/ott-platform/api"
)

func WriteResponse(status int, response interface{}, rw http.ResponseWriter) {
	rw.Header().Add("content-type", "application/json")
	rw.WriteHeader(status)
	if response != nil {
		json.NewEncoder(rw).Encode(response)
	}
}

func WriteErrorResponse(status int, err error, rw http.ResponseWriter) {
	ottErr, ok := err.(*api.OTTError)
	if !ok {
		ottErr = &api.OTTError{
			Code:        0,
			Message:     "failed in serving request",
			Description: err.Error(),
		}
	} else {
		status = ottErr.Code.HTTPStatus()
	}

	rw.Header().Add("content-type", "application/json")
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(ottErr)
}
