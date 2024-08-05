package creational

import "fmt"

// Pizza represents the product that we are building.
type Pizza struct {
	Dough   string
	Sauce   string
	Topping string
}

// PizzaBuilder helps to create a Pizza step by step.
type PizzaBuilder struct {
	dough   string
	sauce   string
	topping string
}

func (b *PizzaBuilder) SetDough(dough string) *PizzaBuilder {
	b.dough = dough
	return b
}

func (b *PizzaBuilder) SetSauce(sauce string) *PizzaBuilder {
	b.sauce = sauce
	return b
}

func (b *PizzaBuilder) SetTopping(topping string) *PizzaBuilder {
	b.topping = topping
	return b
}

func (b *PizzaBuilder) Build() Pizza {
	return Pizza{
		Dough:   b.dough,
		Sauce:   b.sauce,
		Topping: b.topping,
	}
}

func main() {
	// Use the PizzaBuilder to create a Pizza
	builder := &PizzaBuilder{}
	pizza := builder.SetDough("Thin Crust").
		SetSauce("Tomato").
		SetTopping("Cheese").
		Build()

	fmt.Printf("Pizza: %+v\n", pizza)
}
