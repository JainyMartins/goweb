package handler

import (
	"net/http"

	"github.com/JainyMartins/goweb/internal/service"
	"github.com/gin-gonic/gin"
)

// Declaração da Estrutura Request e seus campos rotulados
type request struct {
	Nome        string  `json:"nome"`
	Cor         string  `json:"cor"`
	Preco       float64 `json:"preco"`
	Estoque     int     `json:"estoque"`
	Codigo      string  `json:"codigo"`
	Publicacao  bool    `json:"publicacao"`
	DataCriacao string  `json:"dataCriacao"`
}

// var produtos []repository.Produto

// var lastID int

type Produto struct {
	service service.Service
}

func NewProduct(p service.Service) *Produto {
	return &Produto{
		service: p,
	}
}

/* O método de obtenção de produtos se encarregará de validar a solicitação,
passar a tarefa ao Service e devolver a resposta correspondente ao cliente */

func (c *Produto) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

// Método Salvar
func (c *Produto) Salvar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Validação dos campos obrigatórios
		if req.Nome == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Campo 'nome' é obrigatório"})
			return
		}

		if req.Cor == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Campo 'cor' é obrigatório"})
			return
		}

		if req.Preco <= 0.0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Campo 'preco' é obrigatório"})
			return
		}

		if req.Estoque <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Campo 'estoque' é obrigatório"})
			return
		}

		if req.Codigo == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Campo 'codigo' é obrigatório"})
			return
		}

		if !req.Publicacao {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Campo 'publicacao' é obrigatório"})
			return
		}

		if req.DataCriacao == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Campo 'dataCriacao' é obrigatório"})
			return
		}

		p, err := c.service.Salvar(req.Nome, req.Cor, req.Preco, req.Estoque, req.Codigo, req.Publicacao, req.DataCriacao)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		

		ctx.JSON(http.StatusCreated, p)
	}
}

// func Salvar() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var prod repository.Produto

// 		if err := c.ShouldBindJSON(&prod); err != nil {
// 			c.JSON(http.StatusNotFound, gin.H{"error:": err.Error()})
// 			return
// 		}

// 		//Validação Token
// 		token := c.GetHeader("token")
// 		if token != "123456" {
// 			c.JSON(401, gin.H{
// 				"error": "Você não tem permissão para fazer a solicitação solicitada",
// 			})
// 			return
// 		}

// 		// Validação dos campos obrigatórios
// 		if prod.Nome == "" {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Campo 'nome' é obrigatório"})
// 			return
// 		}

// 		if prod.Cor == "" {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Campo 'cor' é obrigatório"})
// 			return
// 		}

// 		if prod.Preco <= 0.0 {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Campo 'preco' é obrigatório"})
// 			return
// 		}

// 		if prod.Estoque <= 0 {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Campo 'estoque' é obrigatório"})
// 			return
// 		}

// 		if prod.Codigo == "" {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Campo 'codigo' é obrigatório"})
// 			return
// 		}

// 		if !prod.Publicacao {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Campo 'publicacao' é obrigatório"})
// 			return
// 		}

// 		if prod.DataCriacao == "" {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Campo 'dataCriacao' é obrigatório"})
// 			return
// 		}

// 		lastID++
// 		prod.Id = lastID

// 		produtos = append(produtos, prod)
// 		c.JSON(http.StatusCreated, prod)
// 	}
// }

// func Buscar() gin.HandlerFunc{
// 	return func (c *gin.Context){
// 		c.JSON(http.StatusOK, gin.H{
// 			"mensagem": "Olá, Jainy",
// 		})
// 	}
// }

//GetAll com leitura do arquivo produtos.json
// func GetAll(c *gin.Context) {
// 	content, err := os.ReadFile("produtos.json")
// 		if err != nil {
// 			log.Println(err)
// 			c.JSON(500, gin.H{"error": "Falha ao ler o arquivo"})
// 			return
// 		}

// 		var produtos []repository.Produto

// 		err = json.Unmarshal(content, &produtos)

// 		if err != nil {
// 			log.Println(err)
// 			c.JSON(500, gin.H{"error": "Falha ao decodificar os produtos"})
// 			return
// 		}

// 		c.JSON(200, produtos)
// }

//Este GetAll permite a utilização de filtros, para testar na URL tem que colocar após o getAll  ?nomeatributo=nomequequerencontrar
// func GetAllWithFilters(c *gin.Context) {
// 	var produtos []repository.Produto

// 	content, err := os.ReadFile("produtos.json")
// 		if err != nil {
// 			log.Println(err)
// 			c.JSON(500, gin.H{"error": "Falha ao ler o arquivo"})
// 			return
// 		}

// 		err = json.Unmarshal(content, &produtos)

// 		if err != nil {
// 			log.Println(err)
// 			c.JSON(500, gin.H{"error": "Falha ao decodificar os produtos"})
// 			return
// 		}

// 		filteredProdutos := make([]repository.Produto, 0)

// 		//Recebe os valores dos filtros do contexto passados como parâmetro
// 		idFilter, _ := strconv.Atoi(c.Query("id"))
// 		nomeFilter := c.Query("nome")
// 		corFilter := c.Query("cor")
// 		precoFilterStr := c.Query("preco")
// 		estoqueFilter, _ := strconv.Atoi(c.Query("estoque"))
// 		codigoFilter := c.Query("codigo")
// 		publicacaoFilterStr := c.Query("publicacao")
// 		dataCriacaoFilter := c.Query("dataCriacao")

// 		//Convertendo filtro de preço para float64
// 		precoFilter, err := strconv.ParseFloat(precoFilterStr, 64)
// 		if err != nil {
// 			precoFilter = 0.0
// 		}

// 		//Convertendo filtro de publicação para booleano
// 		publicacaoFilter, err := strconv.ParseBool(publicacaoFilterStr)
// 		if err != nil {
// 			publicacaoFilter = false
// 		}

// 		//Gera a lógico do filtro para a matriz de produtos
// 		for _, produto := range produtos {
// 			if idFilter != 0 && produto.Id != idFilter {
// 				continue
// 			}

// 			if nomeFilter != "" && produto.Nome != nomeFilter {
// 				continue
// 			}

// 			if corFilter != "" && produto.Cor != corFilter {
// 				continue
// 			}

// 			if precoFilter != 0.0 && produto.Preco != precoFilter {
// 				continue
// 			}

// 			if estoqueFilter != 0 && produto.Estoque != estoqueFilter {
// 				continue
// 			}

// 			if codigoFilter != "" && produto.Codigo != codigoFilter {
// 				continue
// 			}

// 			if publicacaoFilter && !produto.Publicacao {
// 				continue
// 			}

// 			if dataCriacaoFilter != "" && produto.DataCriacao != dataCriacaoFilter {
// 				continue
// 			}

// 			filteredProdutos = append(filteredProdutos, produto)
// 		}

// 		c.JSON(http.StatusOK, filteredProdutos)
// }

//Função getProductById - aqui vou usar o Handler, que é outra forma que podemos usar para tratar funções, estamos desserializando o array de produtos e buscando por id
//Na URL devemos colocar localhost:8080/produtos/getProductById/numerodoid
// func GetProductById() gin.HandlerFunc{
// 	return func (c *gin.Context){
// 		var produtos []repository.Produto

// 		content, err := os.ReadFile("produtos.json")
// 		if err != nil {
// 			log.Println(err)
// 			c.JSON(500, gin.H{"error": "Falha ao ler o arquivo"})
// 			return
// 		}

// 		err = json.Unmarshal(content, &produtos)

// 		if err != nil {
// 			log.Println(err)
// 			c.JSON(500, gin.H{"error": "Falha ao decodificar os produtos"})
// 			return
// 		}

// 		id, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido."})
// 			return
// 		}

// 		for _, produto := range produtos {
// 			if produto.Id == id {
// 				c.JSON(http.StatusOK, produto)
// 				return
// 			}
// 		}
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Nenhum produto com este ID foi encontrado."})
// 	}
// }
