package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/terehsieh/SeminarioGoLang/internal/config"
	"github.com/terehsieh/SeminarioGoLang/internal/database"
	"github.com/terehsieh/SeminarioGoLang/internal/service/product"
)

func main() {
	//Lee la configuracion
	cfg := readConfig()

	db, err := database.NewDatabase(cfg)
	//cuando termina de ejecutar el main, lanza este
	// defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	//createSchema(db)
	//Inyeccion de la configuracion
	service, _ := product.New(db, cfg) //contructor
	httpService := product.NewHTTPTransport(service)

	//registro el route
	r := gin.Default()
	httpService.Register(r) //endpoints
	r.Run()

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

func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS product (id integer primary key autoincrement, name varchar(100), 
        price integer, description varchar(100));`

	// execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	// or, you can use MustExec, which panics on error
	// insertMessage := `INSERT INTO product (product_name) VALUES (?)`
	// name := fmt.Sprintf("Product number %v")
	// // price := fmt.N("rand.Intn(100)")
	// db.MustExec(insertMessage, name)
	return nil
}

//go run cmd/product/productsrv.go -config ./config/config.yaml
