package pkg

import "fmt"

const (
	ServerType = "server"
	PersnalComputerType = "personal"
	NotebookType = "notebook"
)

type Computer interface {
	GetType() string
	PrintDetails()
}

// есть общая реализация создания компьютера - фабрика
//куда передаем тип в зависимости от которого используем нужные конструкторы
func New(typeName string) Computer {
	switch typeName {
	default: 
		fmt.Printf("%s type does not exist", typeName)
		return nil
	case ServerType:
		return NewServer()
	case NotebookType:
		return NewNotebook()
	case PersnalComputerType:
		return NewPersonalComputer()
	}
}