package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"go-contact-api/handlers"
)

func main() {
	// Configuración del router y las rutas
	router := gin.Default()

	// Endpoints
	router.GET("/contacts", handlers.GetAllContacts)
	router.GET("/contacts/:id", handlers.GetContactByID)
	router.POST("/contacts", handlers.AddContact)
	router.PUT("/contacts/:id", handlers.EditContact)
	router.DELETE("/contacts/:id", handlers.DeleteContact)

	// Iniciar el servidor en el puerto 80
	err := router.Run(":80")
	if err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}

	fmt.Println("Servidor corriendo en http://localhost:80")
}
