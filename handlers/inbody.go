package handlers

import (
	"encoding/json"
	"fmt"
	"gorm/db"
	"gorm/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetInbodys(rw http.ResponseWriter, r *http.Request) {
	inbodys := models.Inbodys{}
	db.Database.Find(&inbodys)
	fmt.Println(inbodys)
	if len(inbodys) == 0 {
		sendError(rw, http.StatusNoContent)
	} else {
		sendData(rw, inbodys, http.StatusOK)
	}
}

func GetInbody(rw http.ResponseWriter, r *http.Request) {
	if inbody, err := GetInbodyById(r); err != nil {
		fmt.Println(err.Error)
		sendError(rw, http.StatusNotFound)
	} else {
		if inbody.IdBody == 0 {
			sendError(rw, http.StatusNoContent)
		} else {
			sendData(rw, inbody, http.StatusOK)
		}
	}
}

func CreateInbody(rw http.ResponseWriter, r *http.Request) {
	inBody := models.InBody{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&inBody); err != nil {
		log.Println(err)
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database.Create(&inBody)
		sendData(rw, inBody, http.StatusOK)
	}
}

func UpdateInBody(rw http.ResponseWriter, r *http.Request) {
	var idInbody int64
	if inBodyAux, err := GetInbodyById(r); err != nil {
		sendError(rw, http.StatusNoContent)
	} else {
		if inBodyAux.IdBody == 0 {
			sendError(rw, http.StatusNoContent)
		} else {
			idInbody = inBodyAux.IdBody
		}

		inbody := models.InBody{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&inbody); err != nil {
			fmt.Println(err)
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			inbody.IdBody = idInbody
			db.Database.Save(&inbody)
			sendData(rw, inbody, http.StatusOK)
		}
	}
}

func DeleteInbody(rw http.ResponseWriter, r *http.Request) {
	if inbody, err := GetInbodyById(r); err != nil {
		fmt.Println(err)
		sendError(rw, http.StatusNotFound)
	} else {
		fmt.Println(inbody)
		db.Database.Delete(&inbody)
		sendData(rw, inbody, http.StatusOK)
	}
}

func GetInbodyById(r *http.Request) (models.InBody, *gorm.DB) {
	vars := mux.Vars(r)
	idInbody, _ := strconv.Atoi(vars["id"])
	inbody := models.InBody{}
	if err := db.Database.First(&inbody, idInbody); err != nil {
		return inbody, err
	}
	fmt.Println(inbody)
	return inbody, nil

}
