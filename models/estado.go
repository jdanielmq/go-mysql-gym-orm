package models

type Estado struct {
	IdEstado    int64  `json:"id_estado"`
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

/*
func NewEstado(descripcion string, habilitado bool) *Estado {
	estado := &Estado{Descripcion: descripcion, Habilitado: habilitado}
	return estado
}

func CreateEstado(descripcion string, habilitado bool) *Estado {
	estado := NewEstado(descripcion, habilitado)
	estado.Save()
	return estado
}


func (estado *Estado) insert() {
	sql := "INSERT INTO db_grossgym_fitness.estado set descripcion=?, habilitado=? "
	result, _ := db.Exec(sql, estado.Descripcion, estado.Habilitado)
	estado.IdEstado, _ = result.LastInsertId()
}

// listar todos los registros
func ListEstados() Estados {
	sql := "SELECT id_estado, descripcion, habilitado FROM db_grossgym_fitness.estado;"
	estados := Estados{}
	rows, _ := db.Query(sql)
	for rows.Next() {
		estado := Estado{}
		rows.Scan(&estado.IdEstado, &estado.Descripcion, &estado.Habilitado)
		estados = append(estados, estado)
	}
	return estados
}

func GetEstado(id int) *Estado {
	estado := NewEstado("", false)
	sql := "SELECT id_estado, descripcion, habilitado FROM db_grossgym_fitness.estado WHERE id_estado=? ;"
	rows, _ := db.Query(sql, id)
	for rows.Next() {
		rows.Scan(&estado.IdEstado, &estado.Descripcion, &estado.Habilitado)
	}
	return estado

}

func (estado *Estado) update() {
	sql := "UPDATE db_grossgym_fitness.estado set descripcion=?, habilitado=? WHERE id_estado = ? ;"
	db.Exec(sql, estado.Descripcion, estado.Habilitado, estado.IdEstado)
}

// guardar o editar un registro
func (estado *Estado) Save() {
	if estado.IdEstado == 0 {
		estado.insert()
	} else {
		estado.update()
	}
}

func (estado *Estado) Delete() {
	sql := "DELETE FROM db_grossgym_fitness.estado WHERE id_estado = ? "
	db.Exec(sql, estado.IdEstado)
}
*/
