package models

import "gorm.io/gorm"

// Album representa un vinilo en la tienda.
type Album struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Year   int     `json:"year"`
	Price  float64 `json:"price"`
}

// Migrate crea la tabla en la base de datos si no existe
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Album{})
}
