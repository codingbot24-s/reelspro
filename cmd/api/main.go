package main

import (
	"log"
	"net/http"
	"reelspro/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main () {
	
	
	r := gin.Default()
	
	r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
			AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
			AllowCredentials: true,
	}))
	
	r.GET("/",handlers.HelloWorldHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}