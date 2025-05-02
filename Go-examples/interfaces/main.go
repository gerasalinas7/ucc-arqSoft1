package main

import "fmt"

// Definimos una interfaz 'Speaker' con el método 'MakeSound'
type Speaker interface {
	MakeSound()
}

// Definimos una estructura 'Cat'
type Cat struct {
	Name string
}

// Definimos una estructura 'Dog'
type Dog struct {
	Name string
}

// Implementación de 'MakeSound' para 'Cat'
func (c Cat) MakeSound() {
	fmt.Println(c.Name + " says: Meow!")
}

// Implementación de 'MakeSound' para 'Dog'
func (d Dog) MakeSound() {
	fmt.Println(d.Name + " says: Woof!")
}

// Función que acepta un 'Speaker' (es decir, cualquier tipo que implemente la interfaz 'Speaker')
func AnimalSound(speaker Speaker) {
	speaker.MakeSound()
}

func main() {
	// Crear instancias de 'Cat' y 'Dog'
	c := Cat{Name: "Whiskers"}
	d := Dog{Name: "Rex"}

	// Parte 1: Usando la interfaz
	// Usamos un slice de 'Speaker' para manejar tanto 'Cat' como 'Dog' de forma uniforme
	animals := []Speaker{c, d}
	for _, animal := range animals {
		AnimalSound(animal)
	}

	// Parte 2: Sin usar la interfaz
	// Aquí trabajamos directamente con las estructuras sin la interfaz
	fmt.Println(c.Name + " says: Meow!") // Código específico para Cat
	fmt.Println(d.Name + " says: Woof!") // Código específico para Dog
}
