package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/src/dto"
)

func Forgot(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var body dto.Forgot
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid Body"})
		return
	}

	if body.OldPassword == "" || body.NewPassword == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "old_password or new_password Is Required"})
		return
	}

	fmt.Println(body)

	fmt.Fprintf(w, "Forgot")
}
