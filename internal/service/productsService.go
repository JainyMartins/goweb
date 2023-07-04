package service

import (
	"context"

	"github.com/JainyMartins/goweb/internal/repository/repositoryutil"
)

type Service interface {
	GetAll() ([]repositoryutil.Product, error)
	Get(id int) (repositoryutil.Product, error)
	Salvar(name, category string, count int, price float64) (repositoryutil.Product, error)
	Update(id int, name, category string, count int, price float64) (repositoryutil.Product, error)
	Delete(id int) error
	UpdateNome(id int, name string) (repositoryutil.Product, error)
	UpdatePreco(id int, preco float64) (repositoryutil.Product, error)
}

type service struct {
	repository repositoryutil.Repository
}

func NewService(r repositoryutil.Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]repositoryutil.Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) Get(id int) (repositoryutil.Product, error) {
	ctx := context.TODO()
	p, err := s.repository.GetOneWithContext(ctx, id)
	if err != nil {
		return repositoryutil.Product{}, err
	}

	return p, nil
}

func (s *service) Salvar(name, category string, count int, price float64) (repositoryutil.Product, error) {
	var product repositoryutil.Product

	product.Name = name
	product.Category = category
	product.Count = count
	product.Price = price

	Product, err := s.repository.Salvar(product)
	if err != nil {
		return repositoryutil.Product{}, err
	}

	return Product, nil
}

func (s *service) Update(id int, name, category string, count int, price float64) (repositoryutil.Product, error) {
	productFind, err := s.Get(id)
	if err != nil {
		return repositoryutil.Product{}, err
	}

	if name != "" {
		productFind.Name = name
	}

	if category != "" {
		productFind.Category = category
	}

	if count != 0 {
		productFind.Count = count
	}

	if price != 0.0 {
		productFind.Price = price
	}

	_, err = s.repository.Update(productFind)
	if err != nil {
		return repositoryutil.Product{}, err
	}
	return productFind, nil

}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) UpdateNome(id int, name string) (repositoryutil.Product, error) {
	product, err := s.repository.UpdateName(id, name)

	return product, err
}

func (s *service) UpdatePreco(id int, price float64) (repositoryutil.Product, error) {
	product, err := s.repository.UpdatePrice(id, price)

	return product, err
}
