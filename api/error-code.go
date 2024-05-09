package api

import "net/http"

type ErrorCode int

const (
	Unknown ErrorCode = iota
	MovieNotFount
	MovieCreationFailed
	MovieUpdateFailed
	UnauthorizedMovie
)

var statusCode = map[ErrorCode]int{
	MovieNotFount: http.StatusNotFound,
}

func (e ErrorCode) HTTPStatus() int {
	if code, ok := statusCode[e]; ok {
		return code
	}
	return http.StatusInternalServerError
}
