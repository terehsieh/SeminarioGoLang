package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HTTPService ...
type HTTPService interface {
	Register(*gin.Engine) // al pasar un router, pone todos los endpoints
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

//Implementa la interfaz HTTPService
type httpService struct {
	endpoints []*endpoint
}

// NewHTTPTransport ...
func NewHTTPTransport(s Service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

//lista de endpoints de httpService
func makeEndpoints(s Service) []*endpoint {
	list := []*endpoint{}

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/products",
		function: getAll(s),
	})

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/products/:id",
		function: getByID(s),
	})

	list = append(list, &endpoint{
		method:   "POST",
		path:     "/products",
		function: create(s),
	})

	list = append(list, &endpoint{
		method:   "PUT",
		path:     "/products/:id",
		function: update(s),
	})

	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/products/:id",
		function: delete(s),
	})

	return list
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"products": s.FindAll(),
		})
	}
}

func getByID(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		i := c.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		f := s.FindByID(id)
		m := http.StatusOK
		if f == nil {
			m = http.StatusNotFound
		}
		c.JSON(m, gin.H{
			"products": f,
		})
	}
}

func create(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p Product
		err := c.BindJSON(&p)
		if err != nil {
			panic(err)
		}
		s.AddProduct(p)
		c.JSON(http.StatusCreated, gin.H{
			"message": "Producto creado",
		})
	}
}

func update(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p Product
		i := c.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		p.ID = id
		err = c.BindJSON(&p)
		if err != nil {
			panic(err)
		}
		s.Update(p)
		c.JSON(http.StatusOK, gin.H{
			"message": "Producto con ID " + i + " modificado",
		})
	}
}

func delete(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		i := c.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		s.Delete(id)
		c.JSON(http.StatusOK, gin.H{
			"message": "ID " + i + " eliminado",
		})
	}
}

// Register ...
//Logica de negocio para ejecutar el handler
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
