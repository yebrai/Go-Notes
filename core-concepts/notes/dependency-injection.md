# Patrón de Inyección de Dependencias en Go

## Descripción
La inyección de dependencias es un patrón que ayuda a hacer que el código sea más modular, fácil de probar y mantener. En lugar de crear dependencias directamente dentro de una función o estructura, se pasan como argumentos.

## Ejemplo
```go
// Definimos una interface para una base de datos
type Database interface {
    Query(query string) string
}

// Implementación de la interface para una base de datos concreta
type MySQLDatabase struct{}

func (db MySQLDatabase) Query(query string) string {
    return "Resultados de la consulta en MySQL"
}

// Función que depende de la interface Database
func GetUserData(db Database, userID int) string {
    return db.Query(fmt.Sprintf("SELECT * FROM users WHERE id = %d", userID))
}

func main() {
    // Inyectar la dependencia
    myDB := MySQLDatabase{}
    userData := GetUserData(myDB, 1)
    fmt.Println(userData)
}
```

En lugar de acoplar la función GetUserData a una base de datos específica, se pasa una interfaz Database, lo que facilita el cambio de implementaciones o las pruebas unitarias.
