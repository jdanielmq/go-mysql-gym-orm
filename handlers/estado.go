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

func GetEstados(rw http.ResponseWriter, r *http.Request) {
	estados := models.Estados{}
	db.Database.Find(&estados)
	fmt.Println(estados)
	if len(estados) == 0 {
		sendError(rw, http.StatusNotFound)
	} else {
		sendData(rw, estados, http.StatusOK)
	}
}

func GetEstado(rw http.ResponseWriter, r *http.Request) {
	if estado, err := getEstadoById(r); err != nil {
		fmt.Println(err.Error)
		sendError(rw, http.StatusNotFound)
	} else {
		if estado.IdEstado == 0 {
			sendError(rw, http.StatusNotFound)
		} else {
			sendData(rw, estado, http.StatusOK)
		}
	}
}

func CreateEstado(rw http.ResponseWriter, r *http.Request) {
	estado := models.Estado{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&estado); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database.Create(&estado)
		sendData(rw, estado, http.StatusCreated)
	}
}

func UpdateEstado(rw http.ResponseWriter, r *http.Request) {
	var idEstado int64
	if estadoAux, err := getEstadoById(r); err != nil {
		sendError(rw, http.StatusNoContent)
	} else {
		if estadoAux.IdEstado == 0 {
			sendError(rw, http.StatusNoContent)
		} else {
			idEstado = estadoAux.IdEstado
		}
		estado := models.Estado{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&estado); err != nil {
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			estado.IdEstado = idEstado
			db.Database.Save(&estado)
			sendData(rw, estado, http.StatusOK)
		}
	}
}

func DeleteEstado(rw http.ResponseWriter, r *http.Request) {
	if estado, err := getEstadoById(r); err != nil {
		fmt.Println(err.Error)
		sendError(rw, http.StatusNotFound)
	} else {
		fmt.Println(estado)
		db.Database.Delete(&estado)
		sendData(rw, estado, http.StatusOK)
	}
}

func getEstadoById(r *http.Request) (models.Estado, *gorm.DB) {
	vars := mux.Vars(r)
	idEstado, _ := strconv.Atoi(vars["id"])
	estado := models.Estado{}
	if err := db.Database.First(&estado, idEstado); err.Error != nil {
		return estado, err
	}
	fmt.Println(estado)
	return estado, nil
}
