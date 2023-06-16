package repository_test

import (
	"errors"
	"testing"

	"github.com/JainyMartins/goweb/internal/repository"
	"github.com/stretchr/testify/assert"
)

type mockStore struct {
	data interface{}
	callingRead bool
}

func (m *mockStore) Read(v interface{}) error {
	m.callingRead = true
	produtos := []repository.Produto{
		{1, "Tomate", "Vermelho", 4.99, 4, "AF333", true, "20231010"},
		{2, "Camarão", "Laranja", 100.99, 5, "AERQW", true, "20230511"},
		{3, "Before Update", "Preto", 50.00, 2, "A123A", true, "20231008"},
	}
	data, ok := v.(*[]repository.Produto)
	if !ok {
		return errors.New("Erro")
	}
	*data = produtos
	return nil
}

func (m *mockStore) Write(v interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	repo := repository.NewRepository(&mockStore{})

	result, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, result, 3)
	assert.Equal(t, 1, result[0].Id)
	assert.Equal(t, "Tomate", result[0].Nome)
	assert.Equal(t, "Vermelho", result[0].Cor)
	assert.Equal(t, 4.99, result[0].Preco)
	assert.Equal(t, 4, result[0].Estoque)
	assert.Equal(t, "AF333", result[0].Codigo)
	assert.Equal(t, true, result[0].Publicacao)
	assert.Equal(t, "20231010", result[0].DataCriacao)

	assert.Equal(t, 2, result[1].Id)
	assert.Equal(t, "Camarão", result[1].Nome)
	assert.Equal(t, "Laranja", result[1].Cor)
	assert.Equal(t, 100.99, result[1].Preco)
	assert.Equal(t, 5, result[1].Estoque)
	assert.Equal(t, "AERQW", result[1].Codigo)
	assert.Equal(t, true, result[1].Publicacao)
	assert.Equal(t, "20230511", result[1].DataCriacao)
}

func TestUpdateNome(t *testing.T){
	mockedStore := &mockStore{}
	repo := repository.NewRepository(mockedStore)

	product := repository.Produto{3, "After Update", "Preto", 50.00, 2, "A123A", true, "20231008"}

	result, err := repo.UpdateNome(3, "After Update")

	assert.NoError(t, err)
	assert.Equal(t, product, result)
	assert.True(t, mockedStore.callingRead)
}