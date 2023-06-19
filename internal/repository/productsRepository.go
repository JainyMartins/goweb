package repository

import (
	"fmt"

	"github.com/JainyMartins/goweb/pkg/store"
)

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

// Armazenamento último ID - não estamos mais usando
// var lastID int

// Criação interface repository com seus métodos
type Repository interface {
	GetAll() ([]Produto, error)
	Get(id int) (Produto, error)
	Salvar(id int, nome, cor string, preco float64, estoque int, codigo string, publicacao bool, dataCriacao string) (Produto, error)
	LastID() (int, error)
	Update(id int, nome, cor string, preco float64, estoque int, codigo string, publicacao bool, dataCriacao string) (Produto, error)
	Delete(id int) error
	UpdateNome(id int, nome string) (Produto, error)
	UpdatePreco(id int, preco float64) (Produto, error)
}

// Criação da estrutura repository para podermos devolver na função de criação de um novo repositório, devolvendo uma interface que deve ter os métodos, nesse caso estamos usando o pacote store para ter acesso aos métodos de escrita e save de arquivo
type repository struct {
	db store.Store
}

// Criação da função que retorna um endereço de memória de uma estrutura de repository que deverá obedecer à interface Repository
// Vamos receber uma store para poder escrever e salvar arquivos json
func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Produto, error) {
	//Aqui fazemos a leitura do array
	if err := r.db.Read(&ps); err != nil {
		return nil, err
	}

	//Aqui retornamos o array de produtos
	return ps, nil
}

func (r *repository) Get(id int) (Produto, error) {
	if err := r.db.Read(&ps); err != nil {
		return Produto{}, err
	}

	for _, produto := range ps {
		if produto.Id == id {
			return produto, nil
		}
	}

	return Produto{}, fmt.Errorf("Produto %d não encontrado", id)
}

func (r *repository) Salvar(id int, nome, cor string, preco float64, estoque int, codigo string, publicacao bool, dataCriacao string) (Produto, error) {
	//Aqui usamos a store para fazermos a leitura do array de produtos
	if err := r.db.Read(&ps); err != nil {
		return Produto{}, err
	}

	//Aqui criamos um produto com os parâmetros passados
	p := Produto{id, nome, cor, preco, estoque, codigo, publicacao, dataCriacao}

	//Aqui add o produto ao array
	ps = append(ps, p)

	//Aqui fazemos a escrita no arquivo e retornamos um erro se houver
	if err := r.db.Write(ps); err != nil {
		return Produto{}, err
	}

	// lastID = p.Id - não usado mais com o pkg store
	//Retornamos o produto salvo
	return p, nil
}

func (r *repository) Update(id int, nome, cor string, preco float64, estoque int, codigo string, publicacao bool, dataCriacao string) (Produto, error) {
	p := Produto{Nome: nome, Cor: cor, Preco: preco, Estoque: estoque, Codigo: codigo, Publicacao: publicacao, DataCriacao: dataCriacao}

	//Aqui usamos a store para fazermos a leitura do array de produtos
	if err := r.db.Read(&ps); err != nil {
		return Produto{}, err
	}

	updated := false
	for i := range ps {
		if ps[i].Id == id {
			p.Id = id
			ps[i] = p
			//Aqui fazemos a escrita no arquivo e retornamos um erro se houver
			if err := r.db.Write(ps[i]); err != nil {
				return Produto{}, err
			}
			updated = true
		}
	}
	if !updated {
		return Produto{}, fmt.Errorf("Produto %d não encontrado", id)
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	//Aqui usamos a store para fazermos a leitura do array de produtos
	if err := r.db.Read(&ps); err != nil {
		return err
	}

	deleted := false
	var index int
	for i := range ps {
		if ps[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("Produto %d não encontrado", id)
	}
	ps = append(ps[:index], ps[index+1:]...)

	//Aqui fazemos a escrita no arquivo e retornamos um erro se houver
	if err := r.db.Write(ps); err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateNome(id int, nome string) (Produto, error) {
	//Aqui usamos a store para fazermos a leitura do array de produtos
	if err := r.db.Read(&ps); err != nil {
		return Produto{}, err
	}

	var p Produto
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			ps[i].Nome = nome
			updated = true
			p = ps[i]
			//Aqui fazemos a escrita no arquivo e retornamos um erro se houver
			if err := r.db.Write(ps[i]); err != nil {
				return Produto{}, err
			}
		}
	}
	if !updated {
		return Produto{}, fmt.Errorf("Produto %d não encontrado", id)
	}
	return p, nil
}

func (r *repository) UpdatePreco(id int, preco float64) (Produto, error) {
	//Aqui usamos a store para fazermos a leitura do array de produtos
	if err := r.db.Read(&ps); err != nil {
		return Produto{}, err
	}

	var p Produto
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			ps[i].Preco = preco
			updated = true
			p = ps[i]
			//Aqui fazemos a escrita no arquivo e retornamos um erro se houver
			if err := r.db.Write(ps[i]); err != nil {
				return Produto{}, err
			}
		}
	}
	if !updated {
		return Produto{}, fmt.Errorf("Produto %d não encontrado", id)
	}
	return p, nil
}

func (r *repository) LastID() (int, error) {
	//Lemos o array de produtos
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}

	//Se não tivermos produtos, vamos ter iniciar no 0
	if len(ps) == 0 {
		return 0, nil
	}

	//Aqui retornamos o último ID de acordo com o tamanho do array de produtos, lembrando que o índice do array começa em 0, por isso fazemos o -1
	return ps[len(ps)-1].Id, nil

	//return lastID, nil - não usado mais com método store
}
