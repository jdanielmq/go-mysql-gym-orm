package models

import "gorm/db"

type Socio struct {
	Rut             string `json:"rut" gorm:"primaryKey"`
	Nombres         string `json:"nombres"`
	ApellidoPaterno string `json:"apellido_paterno"`
	ApellidoMaterno string `json:"apellido_materno"`
	Genero          string `json:"genero"`
	Correo          string `json:"correo"`
	Celular         string `json:"celular"`
	Habilitado      bool   `json:"habilitado"`
}

type Socios []Socio

const SocioShema string = `CREATE TABLE db_grossgym_fitness.socio (
	rut varchar(10) NOT NULL,
	nombres varchar(200) NOT NULL COMMENT 'nombres del socio',
	apellido_paterno varchar(200) NOT NULL,
	apellido_materno varchar(200) NOT NULL,
	genero varchar(100) NOT NULL COMMENT 'Femenino, Masculino.',
	correo varchar(200) NOT NULL,
	celular varchar(100) NOT NULL,
	habilitado BOOL NOT NULL,
	CONSTRAINT socio_pk PRIMARY KEY (rut)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;
`

func MigrarSocios() {
	db.Database.AutoMigrate(Socio{})
}

func (Socio) TableName() string {
	return "socio"
}

/*
func NewSocio(rut string, nombres string, apellidoPaterno string, apellidoMaterno string,
	genero string, correo string, celular string, habilitado bool) *Socio {
	socio := &Socio{Rut: rut, Nombres: nombres, ApellidoPaterno: apellidoPaterno, ApellidoMaterno: apellidoMaterno, Genero: genero, Correo: correo, Celular: celular, Habilitado: habilitado}
	return socio
}

func CreateSocio(rut string, nombres string, apellidoPaterno string, apellidoMaterno string,
	genero string, correo string, celular string, habilitado bool) *Socio {
	socio := NewSocio(rut, nombres, apellidoPaterno, apellidoMaterno, genero, correo, celular, habilitado)
	socio.Save()
	return socio
}

func (socio *Socio) insert() {
	sql := "INSERT INTO db_grossgym_fitness.socio set rut=?, nombres=?, apellido_paterno=?, apellido_materno=?, genero=?, correo=?, celular=?, habilitado=? "
	result, err := db.Exec(sql, socio.Rut, socio.Nombres, socio.ApellidoPaterno, socio.ApellidoMaterno, socio.Genero, socio.Correo, socio.Celular, socio.Habilitado)
	if err != nil {
		log.Fatal(err)
		return
	}
	println(result)
}

// listar todos los registros
func ListSocios() Socios {
	sql := "SELECT rut, nombres, apellido_paterno, apellido_materno, genero, correo, celular, habilitado FROM db_grossgym_fitness.socio;"
	socios := Socios{}
	rows, _ := db.Query(sql)
	for rows.Next() {
		socio := Socio{}
		rows.Scan(&socio.Rut, &socio.Nombres, &socio.ApellidoPaterno, &socio.ApellidoMaterno, &socio.Genero, &socio.Correo,
			&socio.Celular, &socio.Habilitado)
		socios = append(socios, socio)
	}
	return socios
}

func GetSocio(rut string) *Socio {
	socio := NewSocio("", "", "", "", "", "", "", false)
	sql := "SELECT rut, nombres, apellido_paterno, apellido_materno, genero, correo, celular, habilitado FROM db_grossgym_fitness.socio WHERE rut=? ;"
	rows, _ := db.Query(sql, rut)
	for rows.Next() {
		rows.Scan(&socio.Rut, &socio.Nombres, &socio.ApellidoPaterno, &socio.ApellidoMaterno, &socio.Genero, &socio.Correo,
			&socio.Celular, &socio.Habilitado)
	}
	return socio

}

func (socio *Socio) update() {
	sql := "UPDATE db_grossgym_fitness.socio " +
		"SET nombres=?, apellido_paterno=?, apellido_materno=?, genero=?, correo=?, celular=?, habilitado=?  WHERE rut = ? ;"
	db.Exec(sql, socio.Nombres, socio.ApellidoPaterno, socio.ApellidoMaterno, socio.Genero, socio.Correo, socio.Celular, socio.Habilitado, socio.Rut)
}

// guardar o editar un registro
func (socio *Socio) Save() {
	socioAux := GetSocio(socio.Rut)
	if socioAux.Rut == "" {
		socio.insert()
	} else {
		socio.update()
	}
}

func (socio *Socio) Delete() {
	sql := "DELETE FROM db_grossgym_fitness.socio WHERE rut = ? "
	db.Exec(sql, socio.Rut)
}
*/
