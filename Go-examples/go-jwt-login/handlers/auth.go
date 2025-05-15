package handlers

import (
	"net/http"
	"go-jwt-login/utils"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var users = make(map[string]string) // Simula una base de datos en memoria

var mySigningKey = []byte("mysecretkey") // Clave secreta para firmar el JWT

// Estructura para la respuesta de error
type ErrorResponse struct {
	Message string `json:"message"`
}

// Registro de usuario
func Register(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Decodifica el cuerpo de la solicitud
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request"})
		return
	}

	// Hash de la contraseña
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Could not hash password"})
		return
	}

	// Guarda el usuario en el "base de datos"
	users[user.Username] = hashedPassword

	// Responde con un mensaje
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login de usuario
func Login(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Decodifica el cuerpo de la solicitud
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request"})
		return
	}

	// Verifica si el usuario existe
	storedPassword, ok := users[user.Username]
	if !ok {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Invalid username or password"})
		return
	}

	// Compara las contraseñas
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Invalid username or password"})
		return
	}

	// Crea el JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Could not create token"})
		return
	}

	// Responde con el token
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
