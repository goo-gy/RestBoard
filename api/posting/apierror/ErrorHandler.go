package apierror

import (
	"encoding/json"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

func HandleErrorResponse(res http.ResponseWriter, errorResponse ErrorResponse) {
	errorJson, _ := json.Marshal(errorResponse.ErrorBody)
	res.WriteHeader(errorResponse.Status)
	res.Write(errorJson)
}

func BuildErrorResponse(err error) ErrorResponse {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrPostingNotExist()
	}
	return ErrInternalServer()
}
