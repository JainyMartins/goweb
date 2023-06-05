package repository

type Produto struct {
	Id          int     `json:"id"`
	Nome        string  `json:"nome"`
	Cor         string  `json:"cor"`
	Preco       float64 `json:"preco"`
	Estoque     int     `json:"estoque"`
	Codigo      string  `json:"codigo"`
	Publicacao  bool    `json:"publicacao"`
	DataCriacao string  `json:"dataCriacao"`
}

// Armazenamento de produtos
var ps []Produto

// Armazenamento último ID
var lastID int

// Criação interface repository com seus métodos
type Repository interface {
	GetAll() ([]Produto, error)
	Salvar(id int, nome, cor string, preco float64, estoque int, codigo string, publicacao bool, dataCriacao string) (Produto, error)
	LastID() (int, error)
}

// Criação da estrutura repository para podermos devolver na função de criação de um novo repositório, devolvendo uma interface que deve ter os métodos.
type repository struct{}

// Criação da função que retorna um endereço de memória de uma estrutura de repository que deverá obedecer à interface Repository
func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Produto, error) {
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Salvar(id int, nome, cor string, preco float64, estoque int, codigo string, publicacao bool, dataCriacao string) (Produto, error) {
	p := Produto{id, nome, cor, preco, estoque, codigo, publicacao, dataCriacao}
	ps = append(ps, p)
	lastID = p.Id
	return p, nil
}
