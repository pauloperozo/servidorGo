package routes

import (
	"net/http"
	"server/src/controllers"
)

func ProfileRouter(server *http.ServeMux) {

	server.HandleFunc("GET api/profile", controllers.GetProfile)
}
