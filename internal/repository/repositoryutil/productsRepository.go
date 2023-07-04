package repositoryutil

import "context"

type Product struct {
	// Id          int     `json:"id"`
	// Nome        string  `json:"nome"`
	// Cor         string  `json:"cor"`
	// Preco       float64 `json:"preco"`
	// Estoque     int     `json:"estoque"`
	// Codigo      string  `json:"codigo"`
	// Publicacao  bool    `json:"publicacao"`
	// DataCriacao string  `json:"dataCriacao"`
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Count    int     `json:"count"`
	Price    float64 `json:"price"`
}

type ProductFullDataResponse struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Category        string  `json:"category"`
	Count           int     `json:"count"`
	Price           float64 `json:"price"`
	Warehouse       string  `json:"warehouse"`
	WarehouseAdress string  `json:"warehouse_adress"`
}

// Armazenamento de produtos
var ps []Product

// Armazenamento último ID - não estamos mais usando
// var lastID int

// Criação interface repository com seus métodos
type Repository interface {
	GetAll() ([]Product, error)
	Get(id int)(Product, error)
	GetOneWithContext(ctx context.Context, id int) (Product, error)
	Salvar(product Product) (Product, error)
	// LastID() (int, error)
	Update(product Product) (Product, error)
	Delete(id int) error
	UpdateName(id int, name string) (Product, error)
	UpdatePrice(id int, price float64) (Product, error)
}