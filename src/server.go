package server

import (
	"fmt"
	"log"
	"net/http"
	"server/src/routes"
)

func StartServer() {

	server := http.NewServeMux()
	routes.Init(server)

	fmt.Println("Servidor escuchando en el puerto 3000")
	if err := http.ListenAndServe(":3000", server); err != nil {
		log.Fatal(err)
	}
}
