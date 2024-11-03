package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/src/dto"
)

func SignUp(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var body dto.SignUp
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid Body"})
		return
	}

	if body.Email == "" || body.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "email or password Is Required"})
		return
	}

	fmt.Println(body)

	fmt.Fprintf(w, "SignUp")
}
