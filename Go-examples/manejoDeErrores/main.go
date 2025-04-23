package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Definimos una estructura para representar una tarea.
type Task struct {
	Title string
	Done  bool
}

// Creamos un slice para almacenar las tareas.
var tasks []Task

// Función para agregar una tarea.
func addTask(title string) error {
	if title == "" {
		// Error con el uso de errors.New
		return errors.New("el título de la tarea no puede estar vacío")
	}

	// Error wrapping con fmt.Errorf
	tasks = append(tasks, Task{Title: title, Done: false})
	return nil
}

// Función para mostrar todas las tareas.
func listTasks() error {
	if len(tasks) == 0 {
		// Usamos fmt.Errorf para envolver el error con más contexto
		return fmt.Errorf("no hay tareas en la lista")
	}

	fmt.Println("==== Lista de Tareas ====")
	for i, task := range tasks {
		status := "Pendiente"
		if task.Done {
			status = "Completada"
		}
		fmt.Printf("%d. %s - %s\n", i+1, task.Title, status)
	}
	fmt.Println("=========================")
	return nil
}

// Función para marcar una tarea como completada.
func completeTask(index int) error {
	if index < 0 || index >= len(tasks) {
		// Error con error wrapping
		return fmt.Errorf("índice de tarea %d inválido", index)
	}

	tasks[index].Done = true
	return nil
}

// Función para eliminar una tarea.
func deleteTask(index int) error {
	if index < 0 || index >= len(tasks) {
		// Error con error wrapping
		return fmt.Errorf("índice de tarea %d inválido", index)
	}

	// Eliminamos la tarea del slice.
	tasks = append(tasks[:index], tasks[index+1:]...)
	return nil
}

// Función de ejemplo para usar panic y recover.
func panicExample() {
	defer func() {
		// Usamos recover para capturar el panic y evitar que el programa se detenga.
		if r := recover(); r != nil {
			fmt.Println("Error capturado con recover:", r)
		}
	}()

	// Provocamos un panic (ejemplo de error grave).
	panic("¡Algo salió mal! No se pudo completar la operación.")
}

func main() {
	// Usamos defer para garantizar que se cierre el lector de entrada al final.
	defer func() {
		fmt.Println("Cerrando el gestor de tareas...")
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("==== GESTOR DE TAREAS ====")
		fmt.Println("1. Agregar una tarea")
		fmt.Println("2. Listar todas las tareas")
		fmt.Println("3. Marcar tarea como completada")
		fmt.Println("4. Eliminar una tarea")
		fmt.Println("5. Salir")
		fmt.Print("Elige una opción: ")

		// Leer la opción del usuario.
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			// Agregar tarea
			fmt.Print("Ingresa el título de la tarea: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			// Manejar error si el título está vacío.
			if err := addTask(title); err != nil {
				// Usamos fmt.Println para mostrar el error de manera clara.
				fmt.Println("Error al agregar tarea:", err)
			} else {
				fmt.Println("Tarea agregada.")
			}

		case "2":
			// Listar tareas
			if err := listTasks(); err != nil {
				// Usamos fmt.Println para mostrar el error de manera clara.
				fmt.Println("Error al listar tareas:", err)
			}

		case "3":
			// Marcar tarea como completada
			fmt.Print("Ingresa el número de la tarea a completar: ")
			var taskNum int
			_, err := fmt.Scanln(&taskNum)
			if err != nil {
				// Error de tipo al leer un valor inválido
				fmt.Println("Error al leer el número de tarea:", err)
				continue
			}

			if err := completeTask(taskNum - 1); err != nil {
				// Usamos fmt.Println para mostrar el error
				fmt.Println("Error al completar tarea:", err)
			} else {
				fmt.Println("Tarea completada.")
			}

		case "4":
			// Eliminar tarea
			fmt.Print("Ingresa el número de la tarea a eliminar: ")
			var taskNum int
			_, err := fmt.Scanln(&taskNum)
			if err != nil {
				// Error de tipo al leer un valor inválido
				fmt.Println("Error al leer el número de tarea:", err)
				continue
			}

			if err := deleteTask(taskNum - 1); err != nil {
				// Usamos fmt.Println para mostrar el error
				fmt.Println("Error al eliminar tarea:", err)
			} else {
				fmt.Println("Tarea eliminada.")
			}

		case "5":
			// Salir
			fmt.Println("Saliendo del programa...")
			return

		default:
			fmt.Println("Opción no válida. Intenta nuevamente.")
		}

		// Ejemplo de uso de panic y recover
		panicExample()
	}
}
