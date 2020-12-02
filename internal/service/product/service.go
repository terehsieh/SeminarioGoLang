package product

import "github.com/terehsieh/SeminarioGoLang/internal/config"

//Producto,  publica (mayuscula)
type Product struct {
	ID   int64
	Name string
}

//Interfaz, para implementarlo debe implementar todos los metodos
type ProductService interface {
	AddProduct(Product) error
	FindByID(int) *Product
	FindAll() []*Product
}

//privada
// implementa la interfaz ProducService, ya que implementa todos sus metodos
type service struct {
	conf *config.Config
}

//para poder acceder creo el siguiente metodo

//New...
func New(c *config.Config) (ProductService, error) {
	return service{c}, nil
}

func (s service) AddProduct(Product) error {
	return nil
}
func (s service) FindByID(ID int) *Product {
	return nil
}

func (s service) FindAll() []*Product {
	var list []*Product                             // lista de tipo producto
	list = append(list, &Product{0, "Hello World"}) //tiene un puntero (&)
	return list
}
