package main

import (
	"fmt"
	"os"
)

func main() {
	// Abrir un archivo
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// Asegurarse de cerrar el archivo al final de la función main
	defer file.Close()

	// Escribir en el archivo
	file.WriteString("Hello, Go!")
}
