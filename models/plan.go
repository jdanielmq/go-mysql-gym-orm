package models

type Plan struct {
	IdPlan         int64  `json:"id_plan"`
	TipoPlan       string `json:"tipo_plan"`
	MontoPlan      int64  `json:"monto_plan"`
	Duracion       int    `json:"duracion"`
	Unidad         string `json:"unidad"`
	IsMatricula    bool   `json:"isMatricula"`
	MontoMatricula int64  `json:"monto_matricula"`
}

/*
type Planes []Plan

const PlanShema string = `CREATE TABLE db_grossgym_fitness.plan (
	id_plan INT auto_increment NOT NULL,
	tipo_plan varchar(200) NOT NULL COMMENT 'aca el tipo de plan son los siguientes: Diario,Mensual, etc ...',
	monto_plan BIGINT NULL COMMENT 'el precio del plan',
	duracion INT NULL COMMENT 'es la duraccion del plan.',
	unidad varchar(100) NOT NULL COMMENT 'es la unidad de media DIA, MES, AÑO.',
	isMatricula BOOL NOT NULL COMMENT 'es para verificar si tiene matricula',
	monto_matricula BIGINT NOT NULL COMMENT 'valor de la matricula en forma particular',
	CONSTRAINT plan_pk PRIMARY KEY (id_plan)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;
`

func NewPlan(tipoPlan string, montoPlan int64, duracion int, unidad string, isMatricula bool, montoMatricula int64) *Plan {
	plan := &Plan{TipoPlan: tipoPlan, MontoPlan: montoPlan, Duracion: duracion, Unidad: unidad, IsMatricula: isMatricula, MontoMatricula: montoMatricula}
	return plan
}

func CreatePlan(tipoPlan string, montoPlan int64, duracion int, unidad string, isMatricula bool, montoMatricula int64) *Plan {
	plan := NewPlan(tipoPlan, montoPlan, duracion, unidad, isMatricula, montoMatricula)
	plan.Save()
	return plan
}

func (plan *Plan) insert() {
	sql := "INSERT INTO db_grossgym_fitness.plan set tipo_plan=?, monto_plan=?, duracion=?, unidad=?, isMatricula=?, monto_matricula=? "
	result, _ := db.Exec(sql, plan.TipoPlan, plan.MontoPlan, plan.Duracion, plan.Duracion, plan.IsMatricula, plan.MontoMatricula)
	plan.IdPlan, _ = result.LastInsertId()
}

// listar todos los registros
func ListPlanes() Planes {
	sql := "SELECT id_plan, tipo_plan, monto_plan, duracion, unidad, isMatricula, monto_matricula FROM db_grossgym_fitness.plan;"
	planes := Planes{}
	rows, _ := db.Query(sql)
	for rows.Next() {
		plan := Plan{}
		rows.Scan(&plan.IdPlan, &plan.TipoPlan, &plan.MontoPlan, &plan.Duracion, &plan.Unidad, &plan.IsMatricula, &plan.MontoMatricula)
		planes = append(planes, plan)
	}
	return planes
}

func GetPlan(id int) *Plan {
	plan := NewPlan("", 0, 0, "", false, 0)
	sql := "SELECT id_plan, tipo_plan, monto_plan, duracion, unidad, isMatricula, monto_matricula FROM db_grossgym_fitness.plan WHERE id_plan=? ;"
	rows, _ := db.Query(sql, id)
	for rows.Next() {
		rows.Scan(&plan.IdPlan, &plan.TipoPlan, &plan.MontoPlan, &plan.Duracion, &plan.Unidad, &plan.IsMatricula, &plan.MontoMatricula)
	}
	return plan

}

func (plan *Plan) update() {
	sql := "UPDATE db_grossgym_fitness.plan SET tipo_plan=?, monto_plan=?, duracion=?, unidad=?, isMatricula=?, monto_matricula=? WHERE id_plan = ? ;"
	db.Exec(sql, plan.TipoPlan, plan.MontoPlan, plan.Duracion, plan.Duracion, plan.IsMatricula, plan.MontoMatricula, plan.IdPlan)
}

// guardar o editar un registro
func (plan *Plan) Save() {
	if plan.IdPlan == 0 {
		plan.insert()
	} else {
		plan.update()
	}
}

func (plan *Plan) Delete() {
	sql := "DELETE FROM db_grossgym_fitness.plan WHERE id_plan = ? "
	db.Exec(sql, plan.IdPlan)
}
*/
