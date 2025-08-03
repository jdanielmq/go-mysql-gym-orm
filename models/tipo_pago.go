package models

import "gorm/db"

type TipoPago struct {
	IdPago      int64  `json:"id_pago" gorm:"primaryKey"`
	Descripcion string `json:"descripcion"`
	Estado      bool   `json:"estado"`
}
type MediosPagos []TipoPago

func MigrarTiposPagos() {
	db.Database.AutoMigrate(TipoPago{})
}

func (TipoPago) TableName() string {
	return "tipo_pago"
}

/*
const TipoPagoShema string = `CREATE TABLE tipo_pago (
	id_pago INT auto_increment NOT NULL,
	descripcion varchar(200) NOT NULL COMMENT 'aca se debe decribir el tipo de pago (debito, credito, efectivo, etc)',
	estado BOOL NOT NULL COMMENT 'el estado es para habilitar o desactivar',
	CONSTRAINT tipo_pago_pk PRIMARY KEY (id_pago)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;
`

func NewTipoPago(descripcion string, estado bool) *TipoPago {
	tPago := &TipoPago{Descripcion: descripcion, Estado: estado}
	return tPago
}

func CreateTipoPago(descripcion string, estado bool) *TipoPago {
	tpago := NewTipoPago(descripcion, estado)
	tpago.Save()
	return tpago
}

func (tpago *TipoPago) insert() {
	sql := "INSERT db_grossgym_fitness.tipo_pago set descripcion=?, estado=? "
	result, _ := db.Exec(sql, tpago.Descripcion, tpago.Estado)
	tpago.IdPago, _ = result.LastInsertId()
}

// listar todos los registros
func ListTipoPago() (MediosPagos, error) {
	sql := "SELECT id_pago, descripcion, estado FROM db_grossgym_fitness.tipo_pago;"
	tiposPagos := MediosPagos{}
	if rows, err := db.Query(sql); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			tipoPago := TipoPago{}
			rows.Scan(&tipoPago.IdPago, &tipoPago.Descripcion, &tipoPago.Estado)
			tiposPagos = append(tiposPagos, tipoPago)
		}
		return tiposPagos, err
	}

}

func GetTipoPago(id int) (*TipoPago, error) {
	tipoPago := NewTipoPago("", false)
	sql := "SELECT id_pago, descripcion, estado FROM db_grossgym_fitness.tipo_pago WHERE id_pago=? ;"
	if rows, err := db.Query(sql, id); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			rows.Scan(&tipoPago.IdPago, &tipoPago.Descripcion, &tipoPago.Estado)
		}
		return tipoPago, err
	}

}

func (tipoPago *TipoPago) update() {
	sql := "UPDATE db_grossgym_fitness.tipo_pago set descripcion=?, estado=? WHERE id_pago = ? ;"
	db.Exec(sql, tipoPago.Descripcion, tipoPago.Estado, tipoPago.IdPago)

}

// guardar o editar un registro
func (tipoPago *TipoPago) Save() {
	if tipoPago.IdPago == 0 {
		tipoPago.insert()
	} else {
		tipoPago.update()
	}
}

func (tipoPago *TipoPago) Delete() {
	sql := "DELETE FROM db_grossgym_fitness.tipo_pago WHERE id_pago = ? "
	db.Exec(sql, tipoPago.IdPago)
}
*/
