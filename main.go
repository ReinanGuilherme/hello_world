package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")

	router := gin.Default()
	router.GET("/", routerHelloWorld)

	fmt.Println("Server is running on http://localhost:" + os.Getenv("PORT"))
	router.Run("localhost:" + PORT)
}

func routerHelloWorld(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Server On... Hello, World!"})
}
