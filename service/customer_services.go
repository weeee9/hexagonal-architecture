package service

import "github.com/weeee9/hexagonal-architecture/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomerByID(id int) (*domain.Customer, error)
}

func NewCustomerService(repo domain.CustomerRepository) CustomerService {
	return NewDefaultCustomerService(repo)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func NewDefaultCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repo: repo,
	}
}

func (service DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return service.repo.FindAll()
}

func (service DefaultCustomerService) GetCustomerByID(id int) (*domain.Customer, error) {
	return service.repo.GetByID(id)
}
