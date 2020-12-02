package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/terehsieh/SeminarioGoLang/internal/config"
	"github.com/terehsieh/SeminarioGoLang/internal/service/product"
)

func main() {
	//Lee la configuracion
	cfg := readConfig()
	//Inyeccion de la configuracion
	service, _ := product.New(cfg) //contructor
	//Slice de punteros a productos
	for _, m := range service.FindAll() {
		fmt.Println(m)
	}

}

func readConfig() *config.Config {
	// flag: permite leer parametros
	configFile := flag.String("config", "./config.yaml", "this is the service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cfg
}

// func createSchema(db *sqlx.DB) error {
// 	schema := `CREATE TABLE IF NOT EXISTS messages (
// 		id integer primary key autoincrement,
// 		text varchar);`

// 	// execute a query on the server
// 	_, err := db.Exec(schema)
// 	if err != nil {
// 		return err
// 	}

// 	// or, you can use MustExec, which panics on error
// 	insertMessage := `INSERT INTO messages (text) VALUES (?)`
// 	s := fmt.Sprintf("Message number %v", time.Now().Nanosecond())
// 	db.MustExec(insertMessage, s)
// 	return nil
// }
