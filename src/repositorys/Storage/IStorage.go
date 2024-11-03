package repository

type IStorage[T any] interface {
	Create(item T) (T, error)
	Read(id string) (T, error)
	Update(id string, item T) (T, error)
	delete(id string) error
	findAll() ([]T, error)
}
