# Patrón de Factory en Go

## Descripción
El patrón de fábrica se utiliza para crear instancias de objetos sin exponer la lógica de creación al cliente.

## Ejemplo
```go
// Definimos una interface para un animal
type Animal interface {
    Speak() string
}

// Implementaciones de la interface para diferentes animales
type Dog struct{}
func (d Dog) Speak() string { return "Woof!" }

type Cat struct{}
func (c Cat) Speak() string { return "Meow!" }

// Función fábrica que devuelve un objeto de tipo Animal
func AnimalFactory(animalType string) Animal {
    if animalType == "dog" {
        return Dog{}
    } else if animalType == "cat" {
        return Cat{}
    }
    return nil
}

func main() {
    animal := AnimalFactory("dog")
    fmt.Println(animal.Speak())
}
```
Explicación
La función AnimalFactory crea instancias de Dog o Cat según el tipo pasado, escondiendo los detalles de creación.