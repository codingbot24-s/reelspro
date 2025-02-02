package handlers

import (
	"encoding/json"
	"net/http"
)

func UserCreationHandler(w http.ResponseWriter, r *http.Request) {
	// code

	json.NewEncoder(w).Encode(map[string]string{
		"message": "user route working",
	})

	// get the body verify it 


	// hash the password


	// create the user with the given data and struct 
}