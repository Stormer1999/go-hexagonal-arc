package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() CustomerRepository {
	customers := []Customer{
		{CustomerID: 1001, Name: "Ahsh", City: "Bkg", ZipCode: "221C", Status: 0},
		{CustomerID: 1002, Name: "Dss", City: "ASS", ZipCode: "ss2", Status: 1},
	}

	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}
