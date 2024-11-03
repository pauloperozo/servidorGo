package repository

import (
	"errors"
	"server/src/entity"
)

type CustomerRepository struct {
	customers map[string]entity.Customer
}

func New() *CustomerRepository {
	return &CustomerRepository{
		customers: make(map[string]entity.Customer),
	}
}

func (this *CustomerRepository) Create(customer entity.Customer) (entity.Customer, error) {

	return customer, nil

}

func (this *CustomerRepository) Read(id ...string) ([]entity.Customer, error) {

	var allCustomers []entity.Customer
	return allCustomers, nil

}

func (this *CustomerRepository) Update(id string, customer entity.Customer) (entity.Customer, error) {

	return entity.Customer{}, errors.New("customer not found")
}

func (r *CustomerRepository) Delete(id string) (bool, error) {

	return true, nil
}
