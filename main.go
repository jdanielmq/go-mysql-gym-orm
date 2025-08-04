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
	models.MigrarEstados()
	mux := mux.NewRouter()

	/*Endpoints de  los medios de pagos*/
	mux.HandleFunc("/api/tipo-pago/", handlers.GetTiposPagos).Methods("GET")
	mux.HandleFunc("/api/tipo-pago/{id:[0-9]+}", handlers.GetTipoPgo).Methods("GET")
	mux.HandleFunc("/api/tipo-pago/", handlers.CreateTipoPago).Methods("POST")
	mux.HandleFunc("/api/tipo-pago/{id:[0-9]+}", handlers.UpdateTipoPago).Methods("PUT")
	mux.HandleFunc("/api/tipo-pago/{id:[0-9]+}", handlers.DeleteTipoPago).Methods("DELETE")

	/*Endpoints de los estados de las suscripciones*/
	mux.HandleFunc("/api/estado/", handlers.GetEstados).Methods("GET")
	mux.HandleFunc("/api/estado/{id:[0-9]+}", handlers.GetEstado).Methods("GET")
	mux.HandleFunc("/api/estado/", handlers.CreateEstado).Methods("POST")
	mux.HandleFunc("/api/estado/{id:[0-9]+}", handlers.UpdateEstado).Methods("PUT")
	mux.HandleFunc("/api/estado/{id:[0-9]+}", handlers.DeleteEstado).Methods("DELETE")

	/* endpoints de los  */

	log.Fatal(http.ListenAndServe(":3000", mux))

}
