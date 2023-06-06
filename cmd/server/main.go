package main

import (
	"log"

	"github.com/JainyMartins/goweb/cmd/server/handler"
	"github.com/JainyMartins/goweb/internal/repository"
	"github.com/JainyMartins/goweb/internal/service"
	"github.com/JainyMartins/goweb/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

/*
Instanciamos cada camada do domínio Products e usaremos os métodos do controlador para cada endpoint.
*/

// @title MELI Bootcamp API - Produtos
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error ao carregar o arquivo .env")
	}
	
	store := store.Factory("arquivo", "produtos.json")
	if store == nil {
		log.Fatal("Não foi possivel criar a store")
	}

	repo := repository.NewRepository(store)     // Criação da instância Repository
	service := service.NewService(repo) // Criação da instância Service
	p := handler.NewProduct(service)     // Criação do Controller

	r := gin.Default()
	pr := r.Group("/produtos") 
	{
		pr.POST("/post", p.Salvar())
		pr.GET("/getAll", p.GetAll())
		pr.PUT("/:id", p.Update())
		pr.DELETE("/:id", p.Delete())
		pr.PATCH("/updateNome/:id", p.UpdateNome())
		pr.PATCH("/updatePreco/:id", p.UpdatePreco())
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





