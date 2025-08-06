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
	models.MigrarPlanes()
	models.MigrarSocios()
	models.MigrarInstructores()
	models.MigrarInbodys()
	models.MigrarInbodys()

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

	/* endpoints de los planes */
	mux.HandleFunc("/api/plan/", handlers.GetPlanes).Methods("GET")
	mux.HandleFunc("/api/plan/{id:[0-9]+}", handlers.GetPlan).Methods("GET")
	mux.HandleFunc("/api/plan/", handlers.CreatePlan).Methods("POST")
	mux.HandleFunc("/api/plan/{id:[0-9]+}", handlers.UpdatePlan).Methods("PUT")
	mux.HandleFunc("/api/plan/{id:[0-9]+}", handlers.DeletePlan).Methods("DELETE")

	/* endpoints de los socios */
	mux.HandleFunc("/api/socio/", handlers.GetSocios).Methods("GET")
	mux.HandleFunc("/api/socio/{rut:[a-zA-Z0-9-]+}", handlers.GetSocio).Methods("GET")
	mux.HandleFunc("/api/socio/", handlers.CreateSocio).Methods("POST")
	mux.HandleFunc("/api/socio/{rut:[a-zA-Z0-9-]+}", handlers.UpdateSocio).Methods("PUT")
	mux.HandleFunc("/api/socio/{rut:[a-zA-Z0-9-]+}", handlers.DeleteSocio).Methods("DELETE")

	/* endpoints de los instructor */
	mux.HandleFunc("/api/instructor/", handlers.GetInstructores).Methods("GET")
	mux.HandleFunc("/api/instructor/{id:[0-9]+}", handlers.GetInstructor).Methods("GET")
	mux.HandleFunc("/api/instructor/", handlers.CreateInstructor).Methods("POST")
	mux.HandleFunc("/api/instructor/{id:[0-9]+}", handlers.UpdateInstructor).Methods("PUT")
	mux.HandleFunc("/api/instructor/{id:[0-9]+}", handlers.DeleteInstructor).Methods("DELETE")

	/* endpoints de los instructor */
	mux.HandleFunc("/api/inbody/", handlers.GetInbodys).Methods("GET")
	mux.HandleFunc("/api/inbody/{id:[0-9]+}", handlers.GetInbody).Methods("GET")
	mux.HandleFunc("/api/inbody/", handlers.CreateInbody).Methods("POST")
	mux.HandleFunc("/api/inbody/{id:[0-9]+}", handlers.UpdateInBody).Methods("PUT")
	mux.HandleFunc("/api/inbody/{id:[0-9]+}", handlers.DeleteInbody).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", mux))

}
