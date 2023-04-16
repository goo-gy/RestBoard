package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/restBoard/api/domain"
	"github.com/restBoard/api/util"

	"github.com/gorilla/mux"
	"github.com/restBoard/api/apierror"
	"github.com/restBoard/api/service"
)

func HandleCreatePosting(res http.ResponseWriter, req *http.Request) {
	postingRequest := domain.PostingRequest{}
	err := json.NewDecoder(req.Body).Decode(&postingRequest)
	if err != nil {
		apierror.HandleErrorResponse(res, apierror.ErrInvalidInput())
		return
	}
	postingResponse, err := service.CreatePosting(postingRequest)
	if err != nil {
		apierror.HandleErrorResponse(res, apierror.BuildErrorResponse(err))
		return
	}
	util.WriteResponse(res, postingResponse, http.StatusOK)
}

func HandleDetailPosting(res http.ResponseWriter, req *http.Request) {
	pathVariables := mux.Vars(req)
	postingId, err := strconv.ParseInt(pathVariables["postingId"], 10, 64)
	if err != nil {
		apierror.HandleErrorResponse(res, apierror.ErrInvalidInput())
		return
	}

	postingResponse, err := service.DetailPosting(postingId)
	if err != nil {
		apierror.HandleErrorResponse(res, apierror.BuildErrorResponse(err))
		return
	}
	util.WriteResponse(res, postingResponse, http.StatusOK)
}

func HandleListPosting(res http.ResponseWriter, req *http.Request) {
	postingResponseList, err := service.ListPosting()
	if err != nil {
		apierror.HandleErrorResponse(res, apierror.BuildErrorResponse(err))
		return
	}
	util.WriteResponse(res, postingResponseList, http.StatusOK)
}

func HandleUpdatePosting(res http.ResponseWriter, req *http.Request) {
	// Parse Params
	pathVariables := mux.Vars(req)
	postingId, err := strconv.ParseInt(pathVariables["postingId"], 10, 64)
	if err != nil {
		apierror.HandleErrorResponse(res, apierror.ErrInvalidInput())
		return
	}
	postingRequest := domain.PostingRequest{}
	err = json.NewDecoder(req.Body).Decode(&postingRequest)
	if err != nil {
		apierror.HandleErrorResponse(res, apierror.ErrInvalidInput())
		return
	}

	// Do updates
	postingResponse, err := service.UpdatePosting(postingId, postingRequest)
	if err != nil {
		apierror.HandleErrorResponse(res, apierror.BuildErrorResponse(err))
		return
	}

	// Make response
	util.WriteResponse(res, postingResponse, http.StatusOK)
}

func HandleDeletePosting(res http.ResponseWriter, req *http.Request) {
	pathVariables := mux.Vars(req)
	postingId, err := strconv.ParseInt(pathVariables["postingId"], 10, 64)
	if err != nil {
		apierror.HandleErrorResponse(res, apierror.ErrInvalidInput())
		return
	}

	err = service.DeletePosting(postingId)
	if err != nil {
		apierror.HandleErrorResponse(res, apierror.BuildErrorResponse(err))
		return
	}
	util.WriteResponse(res, nil, http.StatusOK)
}
