package apierror

import (
	"net/http"

	"github.com/restBoard/api/util"
)

func HandleErrorResponse(res http.ResponseWriter, err error) {
	errorResponse := errorMap[err.Error()]
	util.WriteResponse(res, errorResponse.ErrorBody, errorResponse.Status)
}