package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)



func ConnectToDB() {
	if err := godotenv.Load() ; err != nil {
		log.Println("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	fmt.Println(uri)

	// connect to mongodb


}