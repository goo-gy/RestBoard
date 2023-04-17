package apierror

import (
	"net/http"
)

type ErrorBody struct {
	Code    string
	Message string
}

type ErrorResponse struct {
	Status    int
	ErrorBody ErrorBody
}

var errorMap = map[string]ErrorResponse {
	"X-500": ErrorResponse {
		Status: http.StatusInternalServerError,
		ErrorBody: ErrorBody{
			Code:    "X-500",
			Message: "Internal Server Error",
		},
	},
	"X-400": ErrorResponse{
		Status: http.StatusBadRequest,
		ErrorBody: ErrorBody{
			Code:    "X-400",
			Message: "Input is Invalid",
		},
	},
	"P-404":ErrorResponse{
		Status: http.StatusNotFound,
		ErrorBody: ErrorBody{
			Code:    "P-404",
			Message: "Posting Not Exist",
		},
	},
}

type ErrInternalServer struct {
}
func (e ErrInternalServer) Error() string {
	return "X-500"
}

type ErrInvalidInput struct {
}
func (e ErrInvalidInput) Error() string {
	return "X-400"
}

type ErrPostingNotExist struct {
}
func (e ErrPostingNotExist) Error() string {
	return "P-404"
}

