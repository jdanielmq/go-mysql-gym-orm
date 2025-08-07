package models

import "gorm/db"

type InBody struct {
	IdBody          int64  `json:"id_body" gorm:"primaryKey"`
	Rut             string `json:"rut" gorm:"type:VARCHAR(10)"`
	Habilitado      bool   `json:"habilitado"`
	JsonInbody      string `json:"json_inbody"`
	FechaEvaluacion string `json:"fecha_evaluacion"`
	Comentario      string `json:"comentario"`
	IdInstructor    int64  `json:"id_instructor"`
}

type Inbodys []InBody

const InBodyShema string = `CREATE TABLE db_grossgym_fitness.inbody (
	id_body INT auto_increment NOT NULL,
	rut varchar(10) NOT NULL,
	habilitado BOOL NOT NULL,
	json_inbody json NOT NULL,
	fecha_evaluacion TIMESTAMP NOT NULL,
	comentario VARCHAR(100) NULL,
	id_instructor INT NULL,
	CONSTRAINT inbody_pk PRIMARY KEY (id_body),
	CONSTRAINT inbody_rut_FK FOREIGN KEY (rut) REFERENCES db_grossgym_fitness.socio(rut),
	CONSTRAINT inbody_instructor_FK FOREIGN KEY (id_instructor) REFERENCES db_grossgym_fitness.instructor(id_instructor)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;
`

func MigrarInbodys() {
	db.Database.AutoMigrate(InBody{})
}

func (InBody) TableName() string {
	return "inbody"
}

/*
func NewInBody(idSocio string, habilitado bool, jsonInbody string, fechaEvaluacion string, comentario string, idInstructor int64) *InBody {
	inbody := &InBody{IdSocio: idSocio, Habilitado: habilitado, JsonInbody: jsonInbody, FechaEvaluacion: fechaEvaluacion, Comentario: comentario, IdInstructor: idInstructor}
	return inbody
}

func CreateInbody(idSocio string, habilitado bool, jsonInbody string, fechaEvaluacion string, comentario string, idInstructor int64) *InBody {
	inbody := NewInBody(idSocio, habilitado, jsonInbody, fechaEvaluacion, comentario, idInstructor)
	inbody.Save()
	return inbody
}

func (inbody *InBody) insert() {
	sql := "INSERT INTO db_grossgym_fitness.inbody  set  id_socio=?, habilitado=?, json_inbody= ?, fecha_evaluacion=?, comentario=?, id_instructor=? ;"
	result, err := db.Exec(sql, inbody.IdSocio, inbody.Habilitado, inbody.JsonInbody, inbody.FechaEvaluacion, inbody.Comentario, inbody.IdInstructor)
	if err != nil {
		log.Println(err)
	}
	inbody.IdBody, _ = result.LastInsertId()
}

// listar todos los registros
func ListInbodys() Inbodys {
	sql := "SELECT id_inbody, id_socio, habilitado, json_inbody, fecha_evaluacion, comentario, id_instructor FROM db_grossgym_fitness.inbody;"
	inbodys := Inbodys{}
	rows, _ := db.Query(sql)
	for rows.Next() {
		inbody := InBody{}
		rows.Scan(&inbody.IdBody, &inbody.IdSocio, &inbody.Habilitado, &inbody.JsonInbody, &inbody.FechaEvaluacion, &inbody.Comentario, &inbody.IdInstructor)
		inbodys = append(inbodys, inbody)
	}
	return inbodys
}

func GetInBody(id int) *InBody {
	inbody := NewInBody("", false, "", "", "", 0)
	sql := "SELECT id_inbody, id_socio, habilitado, json_inbody, fecha_evaluacion, comentario, id_instructor FROM db_grossgym_fitness.inbody WHERE id_inbody=? ;"
	rows, _ := db.Query(sql, id)
	for rows.Next() {
		rows.Scan(&inbody.IdBody, &inbody.IdSocio, &inbody.Habilitado, &inbody.JsonInbody, &inbody.FechaEvaluacion, &inbody.Comentario, &inbody.IdInstructor)
	}
	return inbody

}

func (inbody *InBody) update() {
	sql := "UPDATE db_grossgym_fitness.inbody SET id_socio=?, habilitado=?, json_inbody=?, fecha_evaluacion=?, comentario=?,  id_instructor=? WHERE id_inbody = ? ;"
	db.Exec(sql, inbody.IdSocio, inbody.Habilitado, inbody.JsonInbody, inbody.FechaEvaluacion, inbody.Comentario, inbody.IdInstructor, inbody.IdBody)
}

// guardar o editar un registro
func (inbody *InBody) Save() {
	if inbody.IdBody == 0 {
		inbody.insert()
	} else {
		inbody.update()
	}
}

func (inbody *InBody) Delete() {
	sql := "DELETE FROM db_grossgym_fitness.inbody WHERE id_inbody = ? "
	db.Exec(sql, inbody.IdBody)
}
*/
