package mocks

import (
	"errors"

	"github.com/JainyMartins/goweb/internal/repository"
)

type MockStore struct {
	CallingRead bool
	CallingWrite bool
}

func (m *MockStore) Read(v interface{}) error {
	m.CallingRead = true
	produtos := []repository.Produto{
		{Id: 1, Nome: "Tomate", Cor: "Vermelho", Preco: 4.99, Estoque: 4, Codigo: "AF333", Publicacao: true, DataCriacao: "20231010"},
		{Id: 2, Nome: "Camarão", Cor: "Laranja", Preco: 100.99, Estoque: 5, Codigo: "AERQW", Publicacao: true, DataCriacao: "20230511"},
		{Id: 3, Nome: "Before Update", Cor: "Preto", Preco: 50.00, Estoque: 2, Codigo: "A123A", Publicacao: true, DataCriacao: "20231008"},
	}
	data, ok := v.(*[]repository.Produto)
	if !ok {
		return errors.New("Erro")
	}
	*data = produtos
	return nil
}

func (m *MockStore) Write(v interface{}) error {
	m.CallingWrite = true
	return nil
}