package main

import (
	"log"
	"gorm/config"
	"gorm/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la conexión a la base de datos
	config.InitDB()

	// Migrar las tablas
	config.DB.AutoMigrate()

	// Configuración del router
	router := gin.Default()

	// Endpoints
	router.GET("/albums", handlers.GetAllAlbums)
	router.GET("/albums/:id", handlers.GetAlbumByID)
	router.POST("/albums", handlers.AddAlbum)
	router.PUT("/albums/:id", handlers.UpdateAlbum)
	router.DELETE("/albums/:id", handlers.DeleteAlbum)

	// Iniciar el servidor en el puerto 80
	if err := router.Run(":80"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
