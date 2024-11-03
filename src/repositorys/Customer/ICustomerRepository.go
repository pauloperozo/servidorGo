package repository

import "server/src/entity"

type ICustomerRepository interface {
	Create(customer entity.Customer) (entity.Customer, error)
	Read(id ...string) ([]entity.Customer, error)
	Update(id string, customer entity.Customer) (entity.Customer, error)
	Delete(id string) error
}
