package handler_test

// import (
// 	"bytes"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"fmt"
// 	"testing"

// 	"github.com/JainyMartins/goweb/cmd/server/handler"
// 	"github.com/JainyMartins/goweb/internal/repository"
// 	"github.com/JainyMartins/goweb/internal/service"
// 	"github.com/JainyMartins/goweb/pkg/store"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// func createServer() *gin.Engine {
// 	_ = os.Setenv("TOKEN", "123456")
// 	db := store.Factory("arquivo", "produtos.json")
// 	repo := repository.NewRepository(db)
// 	service := service.NewService(repo)
// 	p := handler.NewProduct(service)
// 	r := gin.Default()

// 	pr := r.Group("/produtos")
// 	pr.POST("/post", p.Salvar())
// 	pr.PUT("/:id", p.Update())
// 	pr.DELETE("/:id", p.Delete())
// 	return r
// }

// func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
// 	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("token", "123456")

// 	return req, httptest.NewRecorder()
// }

// func Test_UpdateProduto_OK(t *testing.T) {
//     r := createServer()
// 	id := "1"
//     // criar uma Request do tipo PUT e Response para obter o resultado
// 	payload := `{"nome": "Uva","cor":"Roxa","preco":10.99,"estoque":10,"codigo":"ARQ12","publicacao":true,"dataCriacao":"20231004"}`
// 	url := fmt.Sprintf("/produtos/%s", id)

//     req, response := createRequestTest(http.MethodPut, url, payload)

//     // diz ao servidor que ele pode atender a solicitação
//     r.ServeHTTP(response, req)

//     assert.Equal(t, 200, response.Code)
// }

// func Test_DeleteProduto_OK(t *testing.T) {
//     r := createServer()
// 	id := "1"
//     // criar uma Request do tipo DELETE e Response para obter o resultado
// 	url := fmt.Sprintf("/produtos/%s", id)

//     req, response := createRequestTest(http.MethodDelete, url, "")

//     // diz ao servidor que ele pode atender a solicitação
//     r.ServeHTTP(response, req)

//     assert.Equal(t, 200, response.Code)
// }

// func Test_SaveProduto_OK(t *testing.T) {
// 	// crie o Servidor e defina as Rotas
// 	r := createServer()
// 	// crie Request do tipo POST e Response para obter o resultado
// 	request, response := createRequestTest(http.MethodPost, "/produtos/post", `{"nome": "Uva","cor":"Roxa","preco":10.99,"estoque":10,"codigo":"ARQ12","publicacao":true,"dataCriacao":"20231004"}`)

// 	// diga ao servidor que ele pode atender a solicitação
// 	r.ServeHTTP(response, request)
// 	assert.Equal(t, 201, response.Code)
// }
