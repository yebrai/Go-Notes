# Patrón de Repositorio en Go

## Descripción
El patrón de repositorio se utiliza para abstraer la lógica de acceso a los datos y aislar las consultas de la capa de dominio.

## Ejemplo
```go
// Definimos un modelo de usuario
type User struct {
    ID   int
    Name string
}

// Interface del repositorio de usuarios
type UserRepository interface {
    GetByID(id int) User
}

// Implementación del repositorio
type InMemoryUserRepository struct {
    users map[int]User
}

func (r *InMemoryUserRepository) GetByID(id int) User {
    return r.users[id]
}

func main() {
    // Inyectamos el repositorio
    repo := &InMemoryUserRepository{
        users: map[int]User{
            1: {ID: 1, Name: "John Doe"},
        },
    }
    
    user := repo.GetByID(1)
    fmt.Printf("Usuario encontrado: %+v\n", user)
}
```

Explicación
InMemoryUserRepository implementa la interfaz UserRepository, permitiendo que la lógica de acceso a datos se maneje de manera abstracta.