package service

import "github.com/JainyMartins/goweb/internal/repository"

type Service interface {
	GetAll() ([]repository.Produto, error)
	Salvar(nome, cor string, preco float64, estoque int, codigo string, publicacao bool, dataCriacao string) (repository.Produto, error)
	Update(id int, nome, cor string, preco float64, estoque int, codigo string, publicacao bool, dataCriacao string) (repository.Produto, error)
	Delete(id int) error
	UpdateNome(id int, name string) (repository.Produto, error)
	UpdatePreco(id int, preco float64) (repository.Produto, error)
}

type service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]repository.Produto, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) Salvar(nome, cor string, preco float64, estoque int, codigo string, publicacao bool, dataCriacao string) (repository.Produto, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return repository.Produto{}, err
	}

	lastID++

	produto, err := s.repository.Salvar(lastID, nome, cor, preco, estoque, codigo, publicacao, dataCriacao)
	if err != nil {
		return repository.Produto{}, err
	}

	return produto, nil
}

func (s *service) Update(id int, nome, cor string, preco float64, estoque int, codigo string, publicacao bool, dataCriacao string) (repository.Produto, error) {

	return s.repository.Update(id, nome, cor, preco, estoque, codigo, publicacao, dataCriacao)
 }

 func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
 }

 func (s *service) UpdateNome(id int, nome string) (repository.Produto, error) {
	return s.repository.UpdateNome(id, nome)
 }

 func (s *service) UpdatePreco(id int, preco float64) (repository.Produto, error) {
	return s.repository.UpdatePreco(id, preco)
 }
