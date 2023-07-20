package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/api/time", getTimeHandler).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
