package util

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(res http.ResponseWriter, responseDto any, status int) {
	responseJson, _ := json.Marshal(responseDto)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	res.Write(responseJson)
}