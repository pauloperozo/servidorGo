package routes

import (
	"net/http"
	"server/src/controllers"
)

func AuthRouter(server *http.ServeMux) {

	server.HandleFunc("POST auth/signup", controllers.SignUp)
	server.HandleFunc("POST auth/signin", controllers.SignIn)
	server.HandleFunc("POST auth/forgot", controllers.Forgot)

}
