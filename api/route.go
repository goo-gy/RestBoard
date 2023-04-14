package api

import (
	"github.com/gorilla/mux"
	"github.com/restBoard/api/posting/controller"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/postings", controller.HandleCreatePosting).Methods("POST")
	router.HandleFunc("/api/postings", controller.HandleListPosting).Methods("GET")
	router.HandleFunc("/api/postings/{postingId}", controller.HandleDetailPosting).Methods("GET")
	router.HandleFunc("/api/postings/{postingId}", controller.HandleUpdatePosting).Methods("PUT")
	router.HandleFunc("/api/postings/{postingId}", controller.HandleDeletePosting).Methods("DELETE")
	return router
}
