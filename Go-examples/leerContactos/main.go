package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// En Go, para que un campo sea exportado (y, por lo tanto, usado por encoding/json),
// debe comenzar con una letra mayúscula.
type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

const pathFile = "contacts.json"

func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create(pathFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Para guardar los datos en formato JSON
	encoder := json.NewEncoder(file) // Prepara para guardar structs como JSON en un archivo
	err = encoder.Encode(contacts)    // Convierte y escribe el data (ej. slice) al archivo
	if err != nil {
		return err
	}

	return nil
}

func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open(pathFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Leer los datos en formato JSON
	decoder := json.NewDecoder(file) // Prepara para leer JSON desde un archivo
	err = decoder.Decode(contacts)    // Lee el JSON del archivo y lo convierte a structs
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var contacts []Contact

	// Intentamos cargar los contactos desde el archivo
	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar los contactos: ", err)
	}

	// Leer desde la entrada estándar (teclado)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(
			"==== GESTOR DE CONTACTOS ====\n"+
				"1. Agregar un contacto \n"+
				"2. Mostrar todos los contactos \n"+
				"3. Salir \n",
		)
		var opcion int
		_, err := fmt.Scanln(&opcion)
		if err != nil {
			fmt.Println("Error al leer la opción: ", err)
			return
		}

		switch opcion {
		case 1:
			// Agregar un contacto
			fmt.Print("Nombre: ")
			nombre, _ := reader.ReadString('\n')
			fmt.Print("Email: ")
			email, _ := reader.ReadString('\n')
			fmt.Print("Teléfono: ")
			telefono, _ := reader.ReadString('\n')

			nombre = strings.TrimSpace(nombre)  // Para quitar \r\n
			email = strings.TrimSpace(email)
			telefono = strings.TrimSpace(telefono)

			con := Contact{Name: nombre, Email: email, Phone: telefono}
			contacts = append(contacts, con)

			// Guardar el contacto en el archivo
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guardar el contacto")
			} else {
				fmt.Println("Contacto guardado exitosamente")
			}

		case 2:
			// Mostrar los contactos
			fmt.Println("=========================")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s | Email: %s | Telefono: %s \n",
					index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("=========================")

		case 3:
			fmt.Println("Saliendo del programa...")
			return

		default:
			fmt.Println("Opción incorrecta.")
		}
	}
}
