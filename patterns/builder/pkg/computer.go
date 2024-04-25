package pkg

import "fmt"

type Computer struct {
	Core int
	Brand string 
	Memory int 
	Monitor int
	GraphicCard int
}

func (pc *Computer) Print() {
	fmt.Printf("%s, Core: %d, Mem: %d, Graphic cad: %d, Monitor: %d\n", pc.Brand, pc.Core, pc.Memory, pc.GraphicCard, pc.Monitor)
}