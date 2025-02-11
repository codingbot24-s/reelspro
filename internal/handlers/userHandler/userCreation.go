package handlers

import (
	"encoding/json"

	"log"
	"net/http"

	"github.com/codingbot24-s/reelspro/internal/database"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func UserCreationHandler(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	var user database.User
	// accept the user data from the request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Fatal("json decode error: ", err)
		http.Error(w, "json decode error", http.StatusBadRequest)
		return
	}
	// validate the user data
 
	validate := validator.New()

	if err := validate.Struct(user); err != nil {
		log.Fatal("validation error: ", err)
		http.Error(w, "validation error", http.StatusBadRequest)
	}

	// hash the password

	hashedpass,err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	user.Password = string(hashedpass)


	if err != nil {
		log.Fatal("hashing error: ", err)
		http.Error(w, "hashing error", http.StatusInternalServerError)
	}

	// create the user with the given data and struct 

	client := database.Client

	userId,err := client.Database("reelspro").Collection("users").InsertOne(r.Context(), user)

	

	if err != nil {
		log.Fatal("user insert error: ", err)
		http.Error(w, "insert error", http.StatusInternalServerError)
		return
	}	

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "user created successfully",
		"user_id": userId.InsertedID,
		"success": "true",
	})

}