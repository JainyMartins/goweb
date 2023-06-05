package main

import (
	"github.com/gin-gonic/gin"
	"github.com/JainyMartins/goweb/internal/service"
	"github.com/JainyMartins/goweb/internal/repository"
	"github.com/JainyMartins/goweb/cmd/server/handler"
)

/*
Instanciamos cada camada do domínio Products e usaremos os métodos do controlador para cada endpoint.
*/

func main() {
	repo := repository.NewRepository()     // Criação da instância Repository
	service := service.NewService(repo) // Criação da instância Service
	p := handler.NewProduct(service)     // Criação do Controller

	r := gin.Default()
	pr := r.Group("/produtos") 
	{
		pr.POST("/post", p.Salvar())
		pr.GET("/getAll", p.GetAll())
	}

	r.Run()
}


//Func main com outros exercícios sem estrutura
// func main() {
// 	r := gin.Default()

// 	pr := r.Group("/produtos")
// 	{
// 		pr.POST("/salvar", handler.)
// 		pr.GET("/getAll", handler.GetAll())
// 		// pr.GET("/get", controller.Buscar())
// 		// pr.GET("/getAll", controller.GetAllWithFilters)
// 		// pr.GET("/getProductById/:id", controller.GetProductById())
// 	}

// 	r.Run()

// }





