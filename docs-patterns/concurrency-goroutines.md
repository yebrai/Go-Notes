# Concurrencia con Goroutines en Go

## Descripción
Las goroutines son funciones o métodos que se ejecutan concurrentemente con otras goroutines dentro del mismo programa.

## Ejemplo
```go
package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hola!")
}

func main() {
    go sayHello() // Ejecuta la función como una goroutine
    time.Sleep(1 * time.Second) // Espera para que la goroutine se ejecute
    fmt.Println("Main terminado")
}
Explicación
La función sayHello se ejecuta concurrentemente en una goroutine, permitiendo que main continúe ejecutándose.