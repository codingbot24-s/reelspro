package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/codingbot24-s/reelspro/internal/database"
	"github.com/codingbot24-s/reelspro/internal/utils"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserResponse struct {
	Message string      `json:"message"`
	UserID  interface{} `json:"user_id,omitempty"`
	Token   string      `json:"token,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func UserCreationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user database.User

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate user data
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Validation error")
		return
	}

	// Check if email already exists
	client := database.Client
	collection := client.Database("reelspro").Collection("users")

	var existingUser database.User
	err := collection.FindOne(r.Context(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		respondWithError(w, http.StatusConflict, "Email already exists")
		return
	}

	// Hash password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error processing request")
		return
	}
	user.Password = string(hashedPass)

	// Insert user
	result, err := collection.InsertOne(r.Context(), user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating user")
		return
	}

	// Generate JWT token
	userID := result.InsertedID.(primitive.ObjectID).Hex()
	token, err := utils.GenerateToken(userID, user.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error generating token")
		return
	}

	// Respond with success
	response := UserResponse{
		Message: "User created successfully",
		UserID:  userID,
		Token:   token,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	response := UserResponse{
		Error: message,
	}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
