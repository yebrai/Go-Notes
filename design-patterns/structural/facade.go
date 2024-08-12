package structural

import "fmt"

// Subsystem 1
type CPU struct{}

func (c *CPU) Start() {
	fmt.Println("CPU started")
}

// Subsystem 2
type Memory struct{}

func (m *Memory) Load() {
	fmt.Println("Memory loaded")
}

// Facade
type ComputerFacade struct {
	cpu    *CPU
	memory *Memory
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{cpu: &CPU{}, memory: &Memory{}}
}

func (cf *ComputerFacade) StartComputer() {
	cf.cpu.Start()
	cf.memory.Load()
}

func main() {
	computer := NewComputerFacade()
	computer.StartComputer()
}
