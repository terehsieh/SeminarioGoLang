package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//CONNECT TO DB
	fmt.Println("Conectando a MySQL: http://localhost:8080")
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Se conecto a la base de MySQL")
	//INSERT
	// insert, err := db.Query("INSERT INTO users VALUES ('TERE')")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()
	// fmt.Println("El usuario fue insertada correctamente")

}
