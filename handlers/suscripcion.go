package handlers

import (
	"encoding/json"
	"fmt"
	"gorm/db"
	"gorm/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetSuscripciones(rw http.ResponseWriter, r *http.Request) {
	suscripciones := models.Suscripciones{}
	db.Database.Find(&suscripciones)
	fmt.Println(suscripciones)
	if len(suscripciones) == 0 {
		sendError(rw, http.StatusNoContent)
	} else {
		sendData(rw, suscripciones, http.StatusOK)
	}
}

func GetSuscripcion(rw http.ResponseWriter, r *http.Request) {
	if suscripcion, err := getSuscripcionById(r); err != nil {
		sendError(rw, http.StatusNoContent)
	} else {
		fmt.Println(suscripcion)
		sendData(rw, suscripcion, http.StatusOK)
	}
}

func CreateSuscripcion(rw http.ResponseWriter, r *http.Request) {
	suscripcion := models.Suscripcion{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&suscripcion); err == nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database.Create(&suscripcion)
		sendData(rw, suscripcion, http.StatusCreated)
	}
}

func UpdateSuscripcion(rw http.ResponseWriter, r *http.Request) {
	var idSuscripcion int64
	if suscripcionAux, err := getSuscripcionById(r); err != nil {
		sendError(rw, http.StatusNoContent)
	} else {
		if suscripcionAux.IdSuscripcion == 0 {
			sendError(rw, http.StatusNotFound)
		} else {
			idSuscripcion = suscripcionAux.IdSuscripcion
		}

		suscripcion := models.Suscripcion{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&suscripcion); err != nil {
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			suscripcion.IdSuscripcion = idSuscripcion
			db.Database.Save(&suscripcion)
			sendData(rw, suscripcion, http.StatusOK)
		}
	}
}

func DeleteSuscripcion(rw http.ResponseWriter, r *http.Request) {
	if suscripcion, err := getSuscripcionById(r); err != nil {
		fmt.Println(err)
		sendError(rw, http.StatusNoContent)
	} else {
		fmt.Println(suscripcion)
		db.Database.Delete(&suscripcion)
		sendData(rw, suscripcion, http.StatusOK)
	}
}

func getSuscripcionById(r *http.Request) (models.Suscripcion, *gorm.DB) {
	vars := mux.Vars(r)
	idSuscripcion, _ := strconv.Atoi(vars["id"])
	suscripscion := models.Suscripcion{}
	if err := db.Database.First(&suscripscion, idSuscripcion); err != nil {
		return suscripscion, err
	}
	fmt.Println(suscripscion)
	return suscripscion, nil
}
