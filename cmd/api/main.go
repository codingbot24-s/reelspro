package main

import (
	"net/http"

	"github.com/codingbot24-s/reelspro/internal/routes/userRoutes"
	"github.com/gorilla/mux"
)



func main () {
	router := mux.NewRouter()

	routes.SetupUserRoutes(router)


	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}

}