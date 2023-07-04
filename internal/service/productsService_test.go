package service_test

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/JainyMartins/goweb/internal/mocks"
// 	"github.com/JainyMartins/goweb/internal/repository"
// 	"github.com/JainyMartins/goweb/internal/service"
// 	"github.com/stretchr/testify/assert"
// )

// // Testando na Camada de Service
// func TestGetAll(t *testing.T) {
// 	repo := repository.NewMySqlRepository(&mocks.MockStore{})
// 	service := service.NewService(repo)

// 	result, err := service.GetAll()
// 	assert.NoError(t, err)
// 	assert.Len(t, result, 3)
// 	assert.Equal(t, 1, result[0].Id)
// 	assert.Equal(t, "Tomate", result[0].Nome)
// 	assert.Equal(t, "Vermelho", result[0].Cor)
// 	assert.Equal(t, 4.99, result[0].Preco)
// 	assert.Equal(t, 4, result[0].Estoque)
// 	assert.Equal(t, "AF333", result[0].Codigo)
// 	assert.Equal(t, true, result[0].Publicacao)
// 	assert.Equal(t, "20231010", result[0].DataCriacao)

// 	assert.Equal(t, 2, result[1].Id)
// 	assert.Equal(t, "Camarão", result[1].Nome)
// 	assert.Equal(t, "Laranja", result[1].Cor)
// 	assert.Equal(t, 100.99, result[1].Preco)
// 	assert.Equal(t, 5, result[1].Estoque)
// 	assert.Equal(t, "AERQW", result[1].Codigo)
// 	assert.Equal(t, true, result[1].Publicacao)
// 	assert.Equal(t, "20230511", result[1].DataCriacao)
// }

// func TestGet(t *testing.T) {
// 	repo := repository.NewRepository(&mocks.MockStore{})
// 	service := service.NewService(repo)

// 	p := repository.Produto{Id: 1, Nome: "Tomate", Cor: "Vermelho", Preco: 4.99, Estoque: 4, Codigo: "AF333", Publicacao: true, DataCriacao: "20231010"}

// 	result, err := service.Get(1)
// 	assert.Equal(t, p, result)
// 	assert.NoError(t, err)
// }

// func TestSalvar(t *testing.T) {
// 	repo := repository.NewRepository(&mocks.MockStore{})
// 	service := service.NewService(repo)

// 	p := repository.Produto{Id: 4, Nome: "Tomate", Cor: "Vermelho", Preco: 4.99, Estoque: 4, Codigo: "AF333", Publicacao: true, DataCriacao: "20231010"}

// 	result, err := service.Salvar(p.Nome, p.Cor, p.Preco, p.Estoque, p.Codigo, p.Publicacao, p.DataCriacao)
// 	assert.Equal(t, p, result)
// 	assert.NoError(t, err)
// }

// func TestGetWithInvalidId(t *testing.T) {
// 	repo := repository.NewRepository(&mocks.MockStore{})
// 	service := service.NewService(repo)

// 	invalidId := 6

// 	_, err := service.Get(invalidId)
// 	assert.EqualError(t, err, fmt.Errorf("Produto %d não encontrado", invalidId).Error())
// }

// func TestUpdate(t *testing.T) {
// 	mockedStore := &mocks.MockStore{}
// 	repo := repository.NewRepository(mockedStore)
// 	service := service.NewService(repo)

// 	product := repository.Produto{Id: 3, Nome: "Laranja", Cor: "Laranja", Preco: 2.99, Estoque: 6, Codigo: "123AE", Publicacao: true, DataCriacao: "20230510"}

// 	result, err := service.Update(3, "Laranja", "Laranja", 2.99, 6, "123AE", true, "20230510")

// 	assert.NoError(t, err)
// 	assert.Equal(t, product, result)
// 	assert.True(t, mockedStore.CallingRead)
// }

// func TestUpdateWithInvalidId(t *testing.T) {
// 	mockedStore := &mocks.MockStore{}
// 	repo := repository.NewRepository(mockedStore)
// 	service := service.NewService(repo)
	
// 	invalidId := 6

// 	_, err := service.Get(invalidId)
// 	assert.EqualError(t, err, fmt.Errorf("Produto %d não encontrado", invalidId).Error())
// }

// func TestUpdateNome(t *testing.T) {
// 	mockedStore := &mocks.MockStore{}
// 	repo := repository.NewRepository(mockedStore)
// 	service := service.NewService(repo)

// 	productName := "After Update"

// 	result, err := service.UpdateNome(3, "After Update")

// 	assert.NoError(t, err)
// 	assert.Equal(t, productName, result.Nome)
// 	assert.True(t, mockedStore.CallingRead)
// }

// func TestUpdatePreco(t *testing.T) {
// 	mockedStore := &mocks.MockStore{}
// 	repo := repository.NewRepository(mockedStore)
// 	service := service.NewService(repo)

// 	productPreco := 10.29

// 	result, err := service.UpdatePreco(3, 10.29)

// 	assert.NoError(t, err)
// 	assert.Equal(t, productPreco, result.Preco)
// 	assert.True(t, mockedStore.CallingRead)
// }

// func TestDelete(t *testing.T) {
// 	mockedStore := &mocks.MockStore{}
// 	repo := repository.NewRepository(mockedStore)
// 	service := service.NewService(repo)

// 	idExists := 3
// 	idDontExists := 6

// 	errExists := service.Delete(idExists)
// 	assert.Nil(t, errExists)

// 	errDontExists := service.Delete(idDontExists)
// 	assert.Error(t, errDontExists)
// 	assert.EqualError(t, errDontExists, fmt.Errorf("Produto %d não encontrado", idDontExists).Error())
// }

