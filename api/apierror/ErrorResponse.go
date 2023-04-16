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

func ErrInternalServer() ErrorResponse {
	return ErrorResponse{
		Status: http.StatusInternalServerError,
		ErrorBody: ErrorBody{
			Code:    "X-500",
			Message: "Internal Server Error",
		},
	}
}

func ErrInvalidInput() ErrorResponse {
	return ErrorResponse{
		Status: http.StatusBadRequest,
		ErrorBody: ErrorBody{
			Code:    "X-400",
			Message: "Input is Invalid",
		},
	}
}

func ErrPostingNotExist() ErrorResponse {
	return ErrorResponse{
		Status: http.StatusNotFound,
		ErrorBody: ErrorBody{
			Code:    "P-404",
			Message: "Posting Not Exist",
		},
	}
}
