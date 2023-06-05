package service

import "github.com/JainyMartins/goweb/internal/repository"

type Service interface {
	GetAll() ([]repository.Produto, error)
	Salvar(nome, cor string, preco float64, estoque int, codigo string, publicacao bool, dataCriacao string) (repository.Produto, error)
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
