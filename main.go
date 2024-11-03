package main

import (
	"fmt"
	repository "server/src/repositorys/Storage"
)

func main() {

	//server.StartServer()

	type Customer struct {
		ID       string
		Email    string
		Password string
	}

	// memory := repository.MemoryRepository[Customer]{
	// 	items: make(map[string]Customer),
	// }

	memory := repository.New[Customer]()

	c := Customer{ID: "123", Email: "P@gmail.com", Password: "hooa"}
	c2 := Customer{ID: "456", Email: "P@gmail.com", Password: "hooa"}
	c3 := Customer{ID: "789", Email: "P@gmail.com", Password: "hooa"}

	memory.Create(c)
	memory.Create(c2)
	memory.Create(c3)

	items, _ := memory.FindAll()

	for _, item := range items {
		fmt.Println(item)
	}

}

/*
repository := MemoryRepository[int]{
    items: make(map[string]int),
}*/
