package behavioral

import "fmt"

// Observer interface
type Observer interface {
	Update(message string)
}

// Concrete observer
type User struct {
	name string
}

func (u *User) Update(message string) {
	fmt.Println(u.name, "received message:", message)
}

// Subject
type Channel struct {
	observers []Observer
}

func (c *Channel) Subscribe(o Observer) {
	c.observers = append(c.observers, o)
}

func (c *Channel) NotifyAll(message string) {
	for _, o := range c.observers {
		o.Update(message)
	}
}

func main() {
	channel := &Channel{}

	user1 := &User{name: "Alice"}
	user2 := &User{name: "Bob"}

	channel.Subscribe(user1)
	channel.Subscribe(user2)

	channel.NotifyAll("New video uploaded")
}
