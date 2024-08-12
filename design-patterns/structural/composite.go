package structural

import "fmt"

// Component interface
type Employee interface {
	ShowDetails()
}

// Leaf struct
type Developer struct {
	name string
}

func (d *Developer) ShowDetails() {
	fmt.Println("Developer:", d.name)
}

// Composite struct
type Manager struct {
	name         string
	subordinates []Employee
}

func (m *Manager) AddSubordinate(e Employee) {
	m.subordinates = append(m.subordinates, e)
}

func (m *Manager) ShowDetails() {
	fmt.Println("Manager:", m.name)
	for _, e := range m.subordinates {
		e.ShowDetails()
	}
}

func main() {
	developer1 := &Developer{name: "Alice"}
	developer2 := &Developer{name: "Bob"}

	manager := &Manager{name: "Charlie"}
	manager.AddSubordinate(developer1)
	manager.AddSubordinate(developer2)

	manager.ShowDetails()
}
