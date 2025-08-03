package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "jdanielmq:JdmQ1481@tcp(localhost:3306)/db_grossgym_fitness?charset=utf8mb4&parseTime=True&loc=Local"
var Database = func() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error en la conexion", err)
		panic(err)
	} else {
		fmt.Println("Conexion Exitosa")
		return db
	}
}()

//const url = "jdanielmq:JdmQ1481@tcp(localhost:3306)/db_grossgym_fitness --->> db_grossgym_orm"
