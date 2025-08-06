package models

import "gorm/db"

type Estado struct {
	IdEstado    int64  `json:"id_estado" gorm:"primaryKey"`
	Descripcion string `json:"descripcion"`
	Habilitado  bool   `json:"habilitado"`
}

type Estados []Estado

const EstadoShema string = `CREATE TABLE db_grossgym_fitness.estado (
	id_estado INT auto_increment NOT NULL,
	descripcion varchar(100) NOT NULL,
	habilitado BOOL NOT NULL,
	CONSTRAINT estado_pk PRIMARY KEY (id_estado)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;
`

func MigrarEstados() {
	db.Database.AutoMigrate(Estado{})
}

func (Estado) TableName() string {
	return "estado"
}
