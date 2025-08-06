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

func GetSocios(rw http.ResponseWriter, r *http.Request) {
	socios := models.Socios{}
	db.Database.Find(&socios)
	fmt.Println(socios)
	if len(socios) == 0 {
		sendError(rw, http.StatusNotFound)
	} else {
		sendData(rw, socios, http.StatusOK)
	}
}

func GetSocio(rw http.ResponseWriter, r *http.Request) {
	if socio, err := getSocioById(r); err != nil {
		fmt.Println(err.Error)
		sendError(rw, http.StatusNotFound)
	} else {
		if socio.Rut == "" {
			sendError(rw, http.StatusNotFound)
		} else {
			sendData(rw, socio, http.StatusOK)
		}
	}
}

func CreateSocio(rw http.ResponseWriter, r *http.Request) {
	socio := models.Socio{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&socio); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		if socio.Rut == "" {
			sendError(rw, http.StatusNoContent)
		} else {
			db.Database.Create(&socio)
			sendData(rw, socio, http.StatusOK)
		}
	}
}

func UpdateSocio(rw http.ResponseWriter, r *http.Request) {
	var rut string
	if socioAux, err := getSocioById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		if socioAux.Rut == "" {
			sendError(rw, http.StatusNotFound)
		} else {
			rut = socioAux.Rut
		}
		socio := models.Socio{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&socio); err != nil {
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			socio.Rut = rut
			db.Database.Save(&socio)
			sendData(rw, socio, http.StatusOK)
		}
	}
}

func DeleteSocio(rw http.ResponseWriter, r *http.Request) {
	if socio, err := getSocioById(r); err != nil {
		fmt.Println(err.Error)
		sendError(rw, http.StatusNotFound)
	} else {
		fmt.Println(socio)
		db.Database.Delete(&socio)
		sendData(rw, socio, http.StatusOK)
	}
}

func getSocioById(r *http.Request) (models.Socio, *gorm.DB) {
	vars := mux.Vars(r)
	idSocio, _ := strconv.Atoi(vars["rut"])
	socio := models.Socio{}
	if err := db.Database.First(&socio, idSocio); err.Error != nil {
		return socio, err
	}
	fmt.Println(socio)
	return socio, nil
}
