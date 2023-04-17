package apierror

import (
	"errors"
	"net/http"

	"github.com/restBoard/api/util"
	"gorm.io/gorm"
)

func HandleErrorResponse(res http.ResponseWriter, errorResponse ErrorResponse) {
	util.WriteResponse(res, errorResponse.ErrorBody, errorResponse.Status)
}

func BuildErrorResponse(err error) ErrorResponse {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrPostingNotExist()
	}
	return ErrInternalServer()
}
