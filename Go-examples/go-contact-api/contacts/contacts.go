package contacts

import "go-contact-api/models"

// Base de datos simulada en memoria.
var Contacts = []models.Contact{
	{ID: "1", Name: "John Doe", Email: "john@example.com", Phone: "123-456-7890"},
	{ID: "2", Name: "Jane Smith", Email: "jane@example.com", Phone: "098-765-4321"},
}

// Función para obtener todos los contactos.
func GetAllContacts() []models.Contact {
	return Contacts
}

// Función para obtener un contacto por su ID.
func GetContactByID(id string) *models.Contact {
	for _, contact := range Contacts {
		if contact.ID == id {
			return &contact
		}
	}
	return nil
}

// Función para agregar un nuevo contacto.
func AddContact(contact models.Contact) {
	Contacts = append(Contacts, contact)
}

// Función para editar un contacto existente.
func EditContact(id string, updatedContact models.Contact) bool {
	for i, contact := range Contacts {
		if contact.ID == id {
			Contacts[i] = updatedContact
			return true
		}
	}
	return false
}

// Función para eliminar un contacto.
func DeleteContact(id string) bool {
	for i, contact := range Contacts {
		if contact.ID == id {
			Contacts = append(Contacts[:i], Contacts[i+1:]...)
			return true
		}
	}
	return false
}
