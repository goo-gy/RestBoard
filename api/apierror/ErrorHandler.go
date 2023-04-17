package apierror

import (
	"net/http"

	"github.com/restBoard/api/util"
)

func HandleErrorResponse(res http.ResponseWriter, err error) {
	errorResponse, exists := errorMap[err.Error()]
	if(exists) {
		util.WriteResponse(res, errorResponse.ErrorBody, errorResponse.Status)
	} else {
		errorResponse = errorMap["X-500"]
		util.WriteResponse(res, errorResponse.ErrorBody, errorResponse.Status)
	}
}