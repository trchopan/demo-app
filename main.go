package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/healthz", getHealthz)

	var addr string
	if os.Getenv("GIN_MODE") == "release" {
		addr = "0.0.0.0:8080"
	} else {
		addr = "localhost:8080"
	}

	router.Run(addr)
}

func getHealthz(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
