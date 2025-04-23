package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"vinyl-store/albums"
	"vinyl-store/models"
	"github.com/google/uuid"
)

// Obtener todos los álbumes
func GetAllAlbums(c *gin.Context) {
	albumList := albums.GetAllAlbums()
	c.JSON(http.StatusOK, albumList)
}

// Obtener un álbum específico por ID
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album := albums.GetAlbumByID(id)
	if album == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Álbum no encontrado"})
		return
	}
	c.JSON(http.StatusOK, album)
}

// Agregar un nuevo álbum
func AddAlbum(c *gin.Context) {
	var newAlbum models.Album
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error al leer datos"})
		return
	}
	newAlbum.ID = uuid.New().String() // Generar un ID único
	albums.AddAlbum(newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

// Actualizar un álbum
func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum models.Album
	if err := c.ShouldBindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error al leer datos"})
		return
	}
	updatedAlbum.ID = id
	if !albums.UpdateAlbum(id, updatedAlbum) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Álbum no encontrado"})
		return
	}
	c.JSON(http.StatusOK, updatedAlbum)
}

// Eliminar un álbum
func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")
	if !albums.DeleteAlbum(id) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Álbum no encontrado"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Álbum eliminado"})
}
