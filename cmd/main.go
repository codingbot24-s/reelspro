package main

import (
	"log"
	"net/http"

	"github.com/codingbot24-s/reelspro/internal/database"
	routes "github.com/codingbot24-s/reelspro/internal/routes/userRoutes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	routes.SetupUserRoutes(router)

	defer database.DisconnectDatabase()

	if err := http.ListenAndServe(":8080", router); err != nil {

		panic(err)
	}
	log.Printf("Server started on port 8080")

}
