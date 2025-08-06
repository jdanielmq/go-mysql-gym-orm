package models

import "gorm/db"

type Suscripcion struct {
	IdSuscripcion  int64  `json:"id_suscripcion" gorm:"primaryKey"`
	FechaCreacion  string `json:"fecha_creacion"`
	NroTransaccion string `json:"nro_transaccion"`
	IdTipoPago     int64  `json:"id_tipo_pago"`
	IdPlan         int64  `json:"id_plan"`
	MontoPlan      int64  `json:"monto_plan"`
	MontoMatricula int64  `json:"monto_matricula"`
	NroCuotas      int    `json:"nro_cuotas"`
	FechaTermino   string `json:"fecha_termino"`
	IdEstado       int64  `json:"id_estado"`
	IdSocio        string `json:"is_socio"`
}

type Suscripciones []Suscripcion

const SuscripcionShema string = `CREATE TABLE db_grossgym_fitness.suscripcion (
	id_suscripcion INT auto_increment NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	nro_transaccion varchar(100) NOT NULL,
	id_tipo_pago INT NOT NULL,
	id_plan INT NOT NULL,
	monto_plan BIGINT NULL,
	monto_matricula BIGINT NULL,
	nro_cuotas INT NULL,
	fecha_termino TIMESTAMP NULL,
	id_estado INT NULL,
	id_socio varchar(10) NOT NULL,
	CONSTRAINT suscripcion_pk PRIMARY KEY (id_suscripcion),
	CONSTRAINT suscripcion_tipo_pago_FK FOREIGN KEY (id_tipo_pago) REFERENCES db_grossgym_fitness.tipo_pago(id_pago),
	CONSTRAINT suscripcion_plan_FK FOREIGN KEY (id_plan) REFERENCES db_grossgym_fitness.plan(id_plan),
	CONSTRAINT suscripcion_estado_FK FOREIGN KEY (id_estado) REFERENCES db_grossgym_fitness.estado(id_estado),
	CONSTRAINT suscripcion_socio_FK FOREIGN KEY (id_socio) REFERENCES db_grossgym_fitness.socio(rut)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;
`

func MigrarSuscripciones() {
	db.Database.AutoMigrate(Suscripcion{})
}

func (Suscripcion) TableName() string {
	return "suscripcion"
}

/*
func NewSuscripcion(fechaCreacion string, nroTransaccion string, idPago int64, idPlan int64, montoPlan int64, montoMatricula int64,
	nroCuotas int, fechaTermino string, idEstado int64, idSocio string) *Suscripcion {
	suscripcion := &Suscripcion{FechaCreacion: fechaCreacion, NroTransaccion: nroTransaccion, IdPago: idPago, IdPlan: idPlan, MontoPlan: montoPlan,
		MontoMatricula: montoMatricula, NroCuotas: nroCuotas, FechaTermino: fechaTermino, IdEstado: idEstado, IdSocio: idSocio}
	return suscripcion
}

func CreateSuscripcion(fechaCreacion string, nroTransaccion string, idPago int64, idPlan int64, montoPlan int64, montoMatricula int64,
	nroCuotas int, fechaTermino string, idEstado int64, idSocio string) *Suscripcion {
	suscripcion := NewSuscripcion(fechaCreacion, nroTransaccion, idPago, idPlan, montoPlan, montoMatricula, nroCuotas, fechaTermino, idEstado, idSocio)
	suscripcion.Save()
	return suscripcion
}

func (suscripcion *Suscripcion) insert() {
	sql := "INSERT INTO db_grossgym_fitness.suscripcion set fecha_creacion=?, nro_transaccion=?, id_tipo_pago=?, id_plan=?, monto_plan=?," +
		" monto_matricula=?, nro_cuotas=?, fecha_termino=?, id_estado=?, id_socio=? "
	result, _ := db.Exec(sql, suscripcion.FechaCreacion, suscripcion.NroTransaccion, suscripcion.IdPago, suscripcion.IdPlan, suscripcion.MontoPlan,
		suscripcion.MontoMatricula, suscripcion.NroCuotas, suscripcion.FechaTermino, suscripcion.IdEstado, suscripcion.IdSocio)
	suscripcion.IdSuscripcion, _ = result.LastInsertId()
}

// listar todos los registros
func ListSuscripciones() Suscripciones {
	sql := "SELECT id_suscripcion, fecha_creacion, nro_transaccion, id_tipo_pago, id_plan, monto_plan, monto_matricula, nro_cuotas, fecha_termino, id_estado, id_socio FROM db_grossgym_fitness.suscripcion;"
	suscripciones := Suscripciones{}
	rows, _ := db.Query(sql)
	for rows.Next() {
		suscripcion := Suscripcion{}
		rows.Scan(&suscripcion.IdSuscripcion, &suscripcion.FechaCreacion, &suscripcion.NroTransaccion, &suscripcion.IdPago, &suscripcion.IdPlan,
			&suscripcion.MontoPlan, &suscripcion.MontoMatricula, &suscripcion.NroCuotas, &suscripcion.FechaTermino, &suscripcion.IdEstado, &suscripcion.IdSocio)
		suscripciones = append(suscripciones, suscripcion)
	}
	return suscripciones
}

func GetSuscripcion(id int) *Suscripcion {
	suscripcion := NewSuscripcion("", "", 0, 0, 0, 0, 0, "", 0, "")
	sql := "SELECT id_suscripcion, fecha_creacion, nro_transaccion, id_tipo_pago, id_plan, monto_plan, monto_matricula, nro_cuotas, fecha_termino, id_estado, id_socio FROM db_grossgym_fitness.suscripcion WHERE id_suscripcion=? ;"
	rows, _ := db.Query(sql, id)
	for rows.Next() {
		rows.Scan(&suscripcion.IdSuscripcion, &suscripcion.FechaCreacion, &suscripcion.NroTransaccion, &suscripcion.IdPago, &suscripcion.IdPlan,
			&suscripcion.MontoPlan, &suscripcion.MontoMatricula, &suscripcion.NroCuotas, &suscripcion.FechaTermino, &suscripcion.IdEstado, &suscripcion.IdSocio)
	}
	return suscripcion

}

func (suscripcion *Suscripcion) update() {
	sql := "UPDATE db_grossgym_fitness.suscripcion SET fecha_creacion=?, nro_transaccion=?, id_tipo_pago=?, id_plan=?, monto_plan=?," +
		" monto_matricula=?, nro_cuotas=?, fecha_termino=?, id_estado=?, id_socio=?  WHERE id_suscripcion = ? ;"
	db.Exec(sql, suscripcion.FechaCreacion, suscripcion.NroTransaccion, suscripcion.IdPago, suscripcion.IdPlan, suscripcion.MontoPlan, suscripcion.MontoMatricula, suscripcion.NroCuotas,
		suscripcion.FechaTermino, suscripcion.IdEstado, suscripcion.IdSocio)
}

// guardar o editar un registro
func (suscripcion *Suscripcion) Save() {
	if suscripcion.IdSuscripcion == 0 {
		suscripcion.insert()
	} else {
		suscripcion.update()
	}
}

func (suscripcion *Suscripcion) Delete() {
	sql := "DELETE FROM db_grossgym_fitness.suscripcion WHERE id_suscripcion = ? "
	db.Exec(sql, suscripcion.IdSuscripcion)
}
*/
