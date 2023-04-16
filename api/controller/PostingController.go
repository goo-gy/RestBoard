package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/restBoard/api/domain"

	"github.com/gorilla/mux"
	"github.com/restBoard/api/apierror"
	"github.com/restBoard/api/service"
)

func HandleCreatePosting(res http.ResponseWriter, req *http.Request) {
	postingRequest := domain.PostingRequest{}
	err := json.NewDecoder(req.Body).Decode(&postingRequest)
	fmt.Println(reflect.TypeOf(err))
	if err != nil {
		apierror.HandleErrorResponse(res, apierror.ErrInvalidInput())
		return
	}
	postingResponse, err := service.CreatePosting(postingRequest)
	if err != nil {
		apierror.HandleErrorResponse(res, apierror.BuildErrorResponse(err))
		return
	}
	createdPostingJson, _ := json.Marshal(postingResponse)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(createdPostingJson)
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
	// TODO : Response 생성 공통화
	createdPostingJson, _ := json.Marshal(postingResponse)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(createdPostingJson)
}

func HandleListPosting(res http.ResponseWriter, req *http.Request) {
	postingResponseList, err := service.ListPosting()
	if err != nil {
		apierror.HandleErrorResponse(res, apierror.BuildErrorResponse(err))
		return
	}
	createdPostingJson, _ := json.Marshal(postingResponseList)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(createdPostingJson)
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
	createdPostingJson, _ := json.Marshal(postingResponse)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(createdPostingJson)
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
	res.WriteHeader(http.StatusOK)
}
