package repository

import (
	"errors"

	"github.com/google/uuid"
)

type MemoryRepository[T any] struct {
	items map[string]T
}

func New[T any]() *MemoryRepository[T] {
	return &MemoryRepository[T]{items: make(map[string]T)}
}

func (this *MemoryRepository[T]) Create(item T) (T, error) {
	id := uuid.New().String()
	this.items[id] = item
	return item, nil
}

func (this *MemoryRepository[T]) Read(id string) (T, error) {

	item, exists := this.items[id]
	if !exists {
		return item, errors.New("item not found")
	}
	return item, nil
}

func (this *MemoryRepository[T]) Update(id string, item T) (T, error) {

	_, exists := this.items[id]
	if !exists {
		return item, errors.New("item not found")
	}
	this.items[id] = item
	return item, nil
}

func (this *MemoryRepository[T]) Delete(id string) error {
	_, exists := this.items[id]
	if !exists {
		return errors.New("item not found")
	}
	delete(this.items, id)
	return nil
}

func (this *MemoryRepository[T]) FindAll() ([]T, error) {
	items := make([]T, 0, len(this.items))
	for _, item := range this.items {
		items = append(items, item)
	}
	return items, nil
}
