package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func createProductTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS product(product_id int primary key auto_increment, product_name text, 
        product_price int, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}
	log.Printf("Rows affected when creating table: %d", rows)
	return nil
}

func main() {
	//CONNECT TO DB
	fmt.Println("Conectando a MySQL: http://localhost:8080")
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Se conecto a la base de MySQL")
	//CREATE TABLE
	err = createProductTable(db)go 
	//INSERT
	// insert, err := db.Query("INSERT INTO users VALUES ('TERE')")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()
	// fmt.Println("El usuario fue insertada correctamente")

}
