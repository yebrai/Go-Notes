package creational

import "fmt"

// Product to be built
type House struct {
	Windows string
	Doors   string
	Roof    string
}

// Builder interface
type HouseBuilder interface {
	SetWindows() HouseBuilder
	SetDoors() HouseBuilder
	SetRoof() HouseBuilder
	Build() House
}

// Concrete builder
type ConcreteHouseBuilder struct {
	house House
}

func (b *ConcreteHouseBuilder) SetWindows() HouseBuilder {
	b.house.Windows = "Wooden Windows"
	return b
}

func (b *ConcreteHouseBuilder) SetDoors() HouseBuilder {
	b.house.Doors = "Wooden Doors"
	return b
}

func (b *ConcreteHouseBuilder) SetRoof() HouseBuilder {
	b.house.Roof = "Concrete Roof"
	return b
}

func (b *ConcreteHouseBuilder) Build() House {
	return b.house
}

func main() {
	builder := &ConcreteHouseBuilder{}
	house := builder.SetWindows().SetDoors().SetRoof().Build()
	fmt.Println(house)
}
