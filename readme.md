# GoNotes

Bienvenido a **GoNotes**. Este repositorio está diseñado para documentar mi experiencia con Go y compartir conceptos. Aquí encontrarás desde fundamentos básicos hasta técnicas avanzadas como patrones de diseño, implementaciones y testing.

## Estructura del Repositorio 📂
```
├── core-concepts # Fundamentos de Go
│ ├── go-by-example # Fundamentos con ejemplos
│ │ └── basics # Conceptos básicos como go-mod, GOPATH, goroutines...
│ └── notes # Notas y documentación adicional
├── design-patterns # Patrones de diseño
│   ├── creational   # Patrones Creacionales
│   ├── structural   # Patrones Estructurales
│   └── behavioral   # Patrones Comportamentales
└── learn_with_examples # Ejemplos prácticos y proyectos en desarrollo
```
El repositorio está organizado en tres secciones principales (por ahora):

### 1. Core Concepts (`core-concepts/`) 
Esta carpeta contiene los fundamentos de Go, organizados en una serie de ejemplos y explicaciones detalladas. Cada archivo dentro de esta carpeta cubre un tema específico y está diseñado para ayudarte a comprender los conceptos básicos del lenguaje.

- **Hello World**
- **Values**
- **Variables**
- **Constants**
- **Control de flujo (If/Else, Switch, For)**
- **Estructuras de Datos (Arrays, Slices, Maps)**
- **Funciones (Funciones, Múltiples Valores de Retorno, Funciones Variádicas)**
- **Otros conceptos importantes (Closures, Recursión, Pointers, Structs, Interfaces, etc.)**
- **Concurrency (Goroutines, Channels, Select, etc.)**

> Cada subtema está documentado con ejemplos de código.

### 2. Design Patterns (`doc-patterns/`) 
Esta carpeta se enfoca en patrones de diseño comunes en Go. Incluye implementaciones prácticas y explicaciones de cómo y cuándo utilizar estos patrones en tus propios proyectos.

- **Singleton**
- **Factory**
- **Observer**
- **Strategy**
- **Decorator**
- **Y más...**

### 3. Learn with Examples (`learn_with_examples/`) 👷‍♂️
Aquí se encuentran ejemplos prácticos donde aplico los conceptos y patrones anteriores, utilizando Test-Driven Development (TDD) para garantizar la calidad y funcionalidad del código.

- **TDD en Go**
- **Ejemplos de proyectos pequeños**
- **Casos de uso específicos**
- **Mejores prácticas**

## Cómo Usar Este Repositorio 🚀

1. **Instalar Go**:
    - Descarga e instala Go desde el [sitio oficial](https://golang.org/dl/).
    - Verifica la instalación ejecutando:
      ```bash
      go version
      ```

2. **Clona o haz fork del repositorio**:
    - Clona el repositorio en tu máquina local utilizando `git`:
      ```bash
      git clone https://github.com/tu_usuario/tu_repositorio.git
      ```
    - Navega a la carpeta del repositorio:
      ```bash
      cd tu_repositorio
      ```

3. **Ejecutar Ejemplos**:
    - Navega a la carpeta que contiene el archivo que quieres ejecutar.
    - Usa el siguiente comando para ejecutar el archivo:
      ```bash
      go run <nombre_del_archivo>.go
      ```
    - Por ejemplo, para ejecutar `hello_world.go` en la carpeta `core_concepts`, usa:
      ```bash
      go run core_concepts/hello_world.go
      ```

---

## Futuras Mejoras 🚀

- **Ampliación de contenidos**: Se agregarán más conceptos avanzados y ejemplos a medida que continúe mi aprendizaje.
- **Mejoras en la documentación**: Planifico añadir más comentarios y explicaciones detalladas en el código para hacer este recurso aún más útil.

## Contribuciones 🤝

Este es un proyecto personal de aprendizaje, pero si tienes sugerencias o encuentras errores, ¡siéntete libre de abrir un issue o enviar un pull request!

### ¿Te Ha Sido Útil? ⭐

Si encuentras útil este repositorio, ¡considera darle una estrella en GitHub! Esto ayuda a que otros también lo descubran.

[¡Dale una estrella!](https://github.com/tu_usuario/tu_repositorio/stargazers)

¡Gracias!

## Referencias 📚

Aquí tienes algunos recursos adicionales de los que voy obteniendo la documentación:

- **Go Route**: [https://roadmap.sh/golang](https://roadmap.sh/golang)
- **Documentación Oficial de Go**: [https://golang.org/doc/](https://golang.org/doc/)
- **Go by Example**: [https://gobyexample.com/](https://gobyexample.com/)
- **Effective Go**: [https://golang.org/doc/effective_go.html](https://golang.org/doc/effective_go.html)
- **Go Patterns**: [https://github.com/tmrts/go-patterns](https://github.com/tmrts/go-patterns)
- **Refactor Guru**: [https://refactoring.guru/](https://refactoring.guru/)
- **Learn Go with Tests**: [https://quii.gitbook.io/learn-go-with-tests](https://quii.gitbook.io/learn-go-with-tests)
