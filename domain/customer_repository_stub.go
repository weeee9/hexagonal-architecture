package domain

import "errors"

func NewCustomerRepotitoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			ID:   1001,
			Name: "Alice",
			City: "New York",
		},
		{
			ID:   1002,
			Name: "Bob",
			City: "Boston",
		},
	}
	return CustomerRepositoryStub{
		customers: customers,
	}
}

type CustomerRepositoryStub struct {
	customers []Customer
}

func (repo CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return repo.customers, nil
}

func (repo CustomerRepositoryStub) GetByID(id int) (*Customer, error) {
	for _, cus := range repo.customers {
		if id == cus.ID {
			return &cus, nil
		}
	}

	return nil, errors.New("customer not found")
}
