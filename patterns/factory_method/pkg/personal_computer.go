package pkg

import "fmt"

type PersnalComputer struct {
	Type string
	Core int
	Memory int
	Monitor bool
}

func NewPersonalComputer() Computer {
	return PersnalComputer{
		Type: PersnalComputerType,
		Core: 8,
		Memory: 16,
		Monitor: true,
	}
}

func (pc PersnalComputer) GetType() string {
	return pc.Type
}

func (pc PersnalComputer) PrintDetails() {
	fmt.Printf("%s Core: %d Mem: %d Monitor: %v\n", pc.Type, pc.Core, pc.Memory, pc.Monitor)
}