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

func GetTiposPagos(rw http.ResponseWriter, r *http.Request) {

	tiposPagos := models.MediosPagos{}
	db.Database.Find(&tiposPagos)
	fmt.Println(tiposPagos)
	sendData(rw, tiposPagos, http.StatusOK)

	/*if tiposPagos, err := models.ListTipoPago(); err != nil {
		models.SendNoFound(rw)
	} else {
		models.SendData(rw, tiposPagos)
	}*/

}

func GetTipoPgo(rw http.ResponseWriter, r *http.Request) {
	if tipoPago, err := getTipoPagoById(r); err != nil {
		fmt.Println(err.Error)
		sendError(rw, http.StatusNotFound)
	} else {
		if tipoPago.IdPago == 0 {
			sendError(rw, http.StatusNotFound)
		} else {
			sendData(rw, tipoPago, http.StatusOK)
		}
	}
}

func getTipoPagoById(r *http.Request) (models.TipoPago, *gorm.DB) {
	vars := mux.Vars(r)
	idTipoPago, _ := strconv.Atoi(vars["id"])
	tipoPago := models.TipoPago{}
	if err := db.Database.First(&tipoPago, idTipoPago); err.Error != nil {
		return tipoPago, err
	}
	fmt.Println(tipoPago)
	return tipoPago, nil
}

func CreateTipoPago(rw http.ResponseWriter, r *http.Request) {
	tipoPago := models.TipoPago{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tipoPago); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database.Create(&tipoPago)
		sendData(rw, tipoPago, http.StatusCreated)
	}
}

func UpdateTipoPago(rw http.ResponseWriter, r *http.Request) {
	var idTipoPago int64
	if tipoPagoAux, err := getTipoPagoById(r); err != nil {
		sendError(rw, http.StatusNoContent)
	} else {
		if tipoPagoAux.IdPago == 0 {
			sendError(rw, http.StatusNoContent)
		} else {
			idTipoPago = tipoPagoAux.IdPago
		}
		tipoPago := models.TipoPago{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&tipoPago); err != nil {
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			tipoPago.IdPago = idTipoPago
			db.Database.Save(&tipoPago)
			sendData(rw, tipoPago, http.StatusOK)
		}
	}

}

func DeleteTipoPago(rw http.ResponseWriter, r *http.Request) {
	if tipoPago, err := getTipoPagoById(r); err != nil {
		fmt.Println(err.Error)
		sendError(rw, http.StatusNotFound)
	} else {
		fmt.Println(tipoPago)
		db.Database.Delete(&tipoPago)
		sendData(rw, tipoPago, http.StatusOK)
	}
}

/*
	if tipoPago, err := getTipoPagoByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		if tipoPago.IdPago == 0 {
			models.SendNoFound(rw)
		} else {
			models.SendData(rw, tipoPago)
		}
	}*/
/*
	rw.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	db.Connect()
	tipoPago, _ := models.GetTipoPago(id)
	db.Close()
	if tipoPago.IdPago == 0 {
		rw.WriteHeader(http.StatusNoContent)
		fmt.Fprintln(rw, "")
	} else {
		output, _ := json.Marshal(tipoPago)
		fmt.Fprintln(rw, string(output))
	}
*/

/*

 */
