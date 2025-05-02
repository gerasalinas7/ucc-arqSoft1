package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm/albums"
	"gorm/models"
	"strconv" // Importamos strconv para convertir el id de string a uint
)

// Obtener todos los álbumes
func GetAllAlbums(c *gin.Context) {
	albumList, err := albums.GetAllAlbums()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al obtener álbumes"})
		return
	}
	c.JSON(http.StatusOK, albumList)
}

// Obtener un álbum específico por ID
func GetAlbumByID(c *gin.Context) {
	// Convertir el ID de string a uint
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)  // Convertir el ID de string a uint
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	album, err := albums.GetAlbumByID(uint(idUint))
	if err != nil {
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
	if err := albums.AddAlbum(newAlbum); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al agregar álbum"})
		return
	}
	c.JSON(http.StatusCreated, newAlbum)
}

// Actualizar un álbum
func UpdateAlbum(c *gin.Context) {
	// Convertir el ID de string a uint
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)  // Convertir el ID de string a uint
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	var updatedAlbum models.Album
	if err := c.ShouldBindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error al leer datos"})
		return
	}
	if err := albums.UpdateAlbum(uint(idUint), updatedAlbum); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Álbum no encontrado"})
		return
	}
	c.JSON(http.StatusOK, updatedAlbum)
}

// Eliminar un álbum
func DeleteAlbum(c *gin.Context) {
	// Convertir el ID de string a uint
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)  // Convertir el ID de string a uint
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}
	if err := albums.DeleteAlbum(uint(idUint)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Álbum no encontrado"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Álbum eliminado"})
}
