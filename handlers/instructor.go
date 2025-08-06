package handlers

import (
	"gorm/db"
	"gorm/models"
	"net/http"
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
