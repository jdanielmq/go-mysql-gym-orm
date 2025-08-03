package main

import (
	"gorm/handlers"
	"gorm/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	models.MigrarTiposPagos()
	mux := mux.NewRouter()
	mux.HandleFunc("/api/tipo-pago/", handlers.GetTiposPagos).Methods("GET")
	mux.HandleFunc("/api/tipo-pago/{id:[0-9]+}", handlers.GetTipoPgo).Methods("GET")
	mux.HandleFunc("/api/tipo-pago/", handlers.CreateTipoPago).Methods("POST")
	mux.HandleFunc("/api/tipo-pago/{id:[0-9]+}", handlers.UpdateTipoPago).Methods("PUT")
	mux.HandleFunc("/api/tipo-pago/{id:[0-9]+}", handlers.DeleteTipoPago).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", mux))

}
