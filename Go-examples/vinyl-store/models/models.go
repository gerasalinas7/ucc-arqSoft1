package models

// Album representa un vinilo en la tienda.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Year   int     `json:"year"`
	Price  float64 `json:"price"`
}
