package app

import (
	"log"
	"muffin_rest/internal/transport/rest"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func Run() {
	router := mux.NewRouter()

	router.Handle("/fin", rest.CheckToken(rest.GetUSDFuncsShares)).Methods("GET")
	router.HandleFunc("/fin", rest.Login).Methods("POST")

	err := http.ListenAndServe(":8080", router); if err!= nil {
		log.Fatal("Error in ListenAndServe: ", err)
	}

}