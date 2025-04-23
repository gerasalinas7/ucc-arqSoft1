package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go-contact-api/contacts"
	"go-contact-api/models"
)

// Obtener todos los contactos.
func GetAllContacts(c *gin.Context) {
	contactsList := contacts.GetAllContacts()
	c.JSON(http.StatusOK, contactsList)
}

// Obtener un contacto por ID.
func GetContactByID(c *gin.Context) {
	id := c.Param("id")
	contact := contacts.GetContactByID(id)
	if contact == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Contacto no encontrado"})
		return
	}
	c.JSON(http.StatusOK, contact)
}

// Agregar un nuevo contacto.
func AddContact(c *gin.Context) {
	var newContact models.Contact
	if err := c.ShouldBindJSON(&newContact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error al leer datos"})
		return
	}
	contacts.AddContact(newContact)
	c.JSON(http.StatusCreated, newContact)
}

// Editar un contacto existente.
func EditContact(c *gin.Context) {
	id := c.Param("id")
	var updatedContact models.Contact
	if err := c.ShouldBindJSON(&updatedContact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error al leer datos"})
		return
	}
	updatedContact.ID = id
	if !contacts.EditContact(id, updatedContact) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Contacto no encontrado"})
		return
	}
	c.JSON(http.StatusOK, updatedContact)
}

// Eliminar un contacto.
func DeleteContact(c *gin.Context) {
	id := c.Param("id")
	if !contacts.DeleteContact(id) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Contacto no encontrado"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contacto eliminado"})
}
