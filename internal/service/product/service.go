package product

import (
	"github.com/jmoiron/sqlx"
	"github.com/terehsieh/SeminarioGoLang/internal/config"
)

//Producto,  publica (mayuscula)
type Product struct {
	ID          int
	Name        string
	Price       int
	Description string
}

func NewProduct(ID int, Name string, Price int,
	Description string) (Product, error) {
	return Product{ID, Name, Price, Description}, nil
}

//Interfaz, para implementarlo debe implementar todos los metodos
type Service interface {
	AddProduct(Product)
	FindByID(int) *Product
	FindAll() []*Product
	Delete(int)
	Update(Product)
}

//privada
// implementa la interfaz ProducService, ya que implementa todos sus metodos
type service struct {
	db   *sqlx.DB
	conf *config.Config
}

//para poder acceder creo el siguiente metodo

//New...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

//CRUD

//CREATE
//Add product
func (s service) AddProduct(p Product) {
	query := `INSERT INTO product (
		 name, price, description) VALUES ( ?, ?, ?)`

	_, err := s.db.Exec(query, p.Name, p.Price, p.Description)

	if err != nil {
		panic(err)
	}
}

//READ
//Get by ID
func (s service) FindByID(ID int) *Product {
	var prod *Product
	if err := s.db.Select(&prod, "SELECT * FROM product WHERE id = ?", ID); err != nil {
		panic(err)
	}
	return prod
}

//Get All
func (s service) FindAll() []*Product {
	var list []*Product                                                 // lista de tipo producto
	if err := s.db.Select(&list, "SELECT * FROM product"); err != nil { // pasa como referencia a list
		panic(err)
	}
	return list
}

//UPDATE
func (s service) Update(p Product) {
	query := `UPDATE product 
		SET name = ?, price = ?, description = ?
		WHERE id = ?`

	_, err := s.db.Exec(query, p.Name,
		p.Price, p.Description, p.ID)

	if err != nil {
		panic(err)
	}
}

// DELETE
func (s service) Delete(ID int) {
	query := "DELETE FROM product WHERE id = ?"

	_, err := s.db.Exec(query, ID)

	if err != nil {
		panic(err)
	}
}
