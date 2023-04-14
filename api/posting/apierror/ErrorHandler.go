package apierror

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

func HandleError(res http.ResponseWriter, errorResponse ErrorResponse) {
	errorJson, _ := json.Marshal(errorResponse.ErrorBody)
	res.WriteHeader(errorResponse.Status)
	res.Write(errorJson)
}

func BuildError(err error) ErrorResponse {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrPostingNotExist()
	}
	return ErrInternalServer()
}
