package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"
	c.JSON(http.StatusOK, resp)
}