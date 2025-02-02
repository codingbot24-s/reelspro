package routes

import (
	

	"github.com/codingbot24-s/reelspro/internal/handlers/userHandler"
	"github.com/gorilla/mux"
)



func SetupUserRoutes (router *mux.Router) {
	
	apiv1 := router.PathPrefix("/api/v1/").Subrouter()
	userRouter := apiv1.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/create", handlers.UserCreationHandler).Methods("GET")
}

