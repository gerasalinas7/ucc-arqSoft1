package main

import (
	"fmt"
	"log"
	"go-jwt-login/handlers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	// Inicializa Gin
	router := gin.Default()

	router.Use(cors.Default())
	// Rutas
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	// Arrancar el servidor
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
