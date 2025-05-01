package albums

import (
	"gorm/models"
	"gorm/config"
)

// Obtener todos los álbumes de la base de datos
func GetAllAlbums() ([]models.Album, error) {
	var albums []models.Album
	result := config.DB.Find(&albums)
	return albums, result.Error
}

// Obtener un álbum por su ID
func GetAlbumByID(id uint) (*models.Album, error) {
	var album models.Album
	result := config.DB.First(&album, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &album, nil
}

// Agregar un nuevo álbum
func AddAlbum(album models.Album) error {
	result := config.DB.Create(&album)
	return result.Error
}

// Actualizar un álbum existente
func UpdateAlbum(id uint, updatedAlbum models.Album) error {
	var album models.Album
	result := config.DB.First(&album, id)
	if result.Error != nil {
		return result.Error
	}
	// Actualizar los campos
	album.Title = updatedAlbum.Title
	album.Artist = updatedAlbum.Artist
	album.Year = updatedAlbum.Year
	album.Price = updatedAlbum.Price
	return config.DB.Save(&album).Error
}

// Eliminar un álbum
func DeleteAlbum(id uint) error {
	result := config.DB.Delete(&models.Album{}, id)
	return result.Error
}
