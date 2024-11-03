package routes

import (
	"net/http"
)

func Init(server *http.ServeMux) {

	AuthRouter(server)
	ProfileRouter(server)

}
