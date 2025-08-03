package models

type Instructor struct {
	IdInstructor     int64  `json:"id_instructor"`
	NombreInstructor string `json:"nombre_instructor"`
	Habilitado       bool   `json:"habilitado"`
}

/*
type Instructores []Instructor

const InstructorShema string = `CREATE TABLE db_grossgym_fitness.instructor (
	id_instructor INT auto_increment NOT NULL,
	nombre_instructor varchar(200) NOT NULL,
	habilitado BOOL NOT NULL,
	CONSTRAINT instructor_pk PRIMARY KEY (id_instructor)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;
`

func NewInstructor(nombreInstructor string, habilitado bool) *Instructor {
	instructor := &Instructor{NombreInstructor: nombreInstructor, Habilitado: habilitado}
	return instructor
}

func CreateInstructor(nombreInstructor string, habilitado bool) *Instructor {
	instructor := NewInstructor(nombreInstructor, habilitado)
	instructor.Save()
	return instructor
}

func (instructor *Instructor) insert() {
	sql := "INSERT INTO db_grossgym_fitness.instructor set nombre_instructor=?, habilitado=? "
	result, _ := db.Exec(sql, instructor.NombreInstructor, instructor.Habilitado)
	instructor.IdInstructor, _ = result.LastInsertId()
}

// listar todos los registros
func ListInstructores() Instructores {
	sql := "SELECT id_instructor, nombre_instructor, habilitado FROM db_grossgym_fitness.instructor;"
	instructores := Instructores{}
	rows, _ := db.Query(sql)
	for rows.Next() {
		instructor := Instructor{}
		rows.Scan(&instructor.IdInstructor, &instructor.NombreInstructor, &instructor.Habilitado)
		instructores = append(instructores, instructor)
	}
	return instructores
}

func GetInstructor(id int) *Instructor {
	instructor := NewInstructor("", false)
	sql := "SELECT id_instructor, nombre_instructor, habilitado FROM db_grossgym_fitness.instructor WHERE id_instructor=? ;"
	rows, _ := db.Query(sql, id)
	for rows.Next() {
		rows.Scan(&instructor.IdInstructor, &instructor.NombreInstructor, &instructor.Habilitado)
	}
	return instructor

}

func (instructor *Instructor) update() {
	sql := "UPDATE db_grossgym_fitness.instructor SET nombre_instructor=?, habilitado=? WHERE id_instructor = ? ;"
	db.Exec(sql, instructor.NombreInstructor, instructor.Habilitado, instructor.IdInstructor)
}

// guardar o editar un registro
func (instructor *Instructor) Save() {
	if instructor.IdInstructor == 0 {
		instructor.insert()
	} else {
		instructor.update()
	}
}

func (instructor *Instructor) Delete() {
	sql := "DELETE FROM db_grossgym_fitness.instructor WHERE id_instructor = ? "
	db.Exec(sql, instructor.IdInstructor)
}
*/
