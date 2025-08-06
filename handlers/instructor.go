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

func GetInstructores(rw http.ResponseWriter, r *http.Request) {
	instructores := models.Instructores{}
	db.Database.Find(&instructores)
	if len(instructores) == 0 {
		sendError(rw, http.StatusNotFound)
	} else {
		sendData(rw, instructores, http.StatusOK)
	}
}

func GetInstructor(rw http.ResponseWriter, r *http.Request) {
	if instructor, err := getInstructorById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		if instructor.IdInstructor == 0 {
			sendError(rw, http.StatusNoContent)
		} else {
			sendData(rw, instructor, http.StatusOK)
		}
	}
}

func CreateInstructor(rw http.ResponseWriter, r *http.Request) {
	instructor := models.Instructor{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&instructor); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database.Create(&instructor)
		sendData(rw, instructor, http.StatusOK)
	}
}

func UpdateInstructor(rw http.ResponseWriter, r *http.Request) {
	var idInstructor int64
	if instructorAux, err := getInstructorById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		if instructorAux.IdInstructor == 0 {
			sendError(rw, http.StatusNoContent)
		} else {
			idInstructor = instructorAux.IdInstructor
		}

		instructor := models.Instructor{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&instructor); err != nil {
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			instructor.IdInstructor = idInstructor
			db.Database.Save(&instructor)
			sendData(rw, instructor, http.StatusOK)
		}
	}
}

func DeleteInstructor(rw http.ResponseWriter, r *http.Request) {
	if instructor, err := getInstructorById(r); err != nil {
		log.Println(err.Error)
		sendError(rw, http.StatusNotFound)
	} else {
		fmt.Println(instructor)
		db.Database.Delete(&instructor)
		sendData(rw, instructor, http.StatusOK)
	}
}

func getInstructorById(r *http.Request) (models.Instructor, *gorm.DB) {

	vars := mux.Vars(r)
	idInstructor, _ := strconv.Atoi(vars["id"])
	instructor := models.Instructor{}
	if err := db.Database.First(&instructor, idInstructor); err != nil {
		return instructor, err
	}
	fmt.Println(instructor)
	return instructor, nil
}
