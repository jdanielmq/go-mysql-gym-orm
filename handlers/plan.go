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

func GetPlanes(rw http.ResponseWriter, r *http.Request) {
	planes := models.Planes{}
	db.Database.Find(&planes)
	fmt.Println(planes)
	if len(planes) == 0 {
		sendError(rw, http.StatusNotFound)
	} else {
		sendData(rw, planes, http.StatusOK)
	}
}

func GetPlan(rw http.ResponseWriter, r *http.Request) {
	if plan, err := getPlanById(r); err != nil {
		fmt.Println(err.Error)
		sendError(rw, http.StatusNotFound)
	} else {
		if plan.IdPlan == 0 {
			sendError(rw, http.StatusNotFound)
		} else {
			sendData(rw, plan, http.StatusOK)
		}
	}
}

func CreatePlan(rw http.ResponseWriter, r *http.Request) {
	plan := models.Plan{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&plan); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database.Create(&plan)
		sendData(rw, plan, http.StatusCreated)
	}
}

func UpdatePlan(rw http.ResponseWriter, r *http.Request) {
	var idPlan int64
	if planAux, err := getPlanById(r); err != nil {
		sendError(rw, http.StatusNoContent)
	} else {
		if planAux.IdPlan == 0 {
			sendError(rw, http.StatusNoContent)
		} else {
			idPlan = planAux.IdPlan
		}
		plan := models.Plan{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&plan); err != nil {
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			plan.IdPlan = idPlan
			db.Database.Save(&plan)
			sendData(rw, plan, http.StatusOK)
		}
	}
}

func DeletePlan(rw http.ResponseWriter, r *http.Request) {
	if plan, err := getPlanById(r); err != nil {
		fmt.Println(err.Error)
		sendError(rw, http.StatusNotFound)
	} else {
		fmt.Println(plan)
		db.Database.Delete(&plan)
		sendData(rw, plan, http.StatusOK)
	}
}

func getPlanById(r *http.Request) (models.Plan, *gorm.DB) {
	vars := mux.Vars(r)
	idPlan, _ := strconv.Atoi(vars["id"])
	plan := models.Plan{}
	if err := db.Database.First(&plan, idPlan); err.Error != nil {
		return plan, err
	}
	fmt.Println(plan)
	return plan, nil
}
