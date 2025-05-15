package main

import (
	"fmt"
	"log"
	"go-jwt-login/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa Gin
	router := gin.Default()

	// Rutas
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	// Arrancar el servidor
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
