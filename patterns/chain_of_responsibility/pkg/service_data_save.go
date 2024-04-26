package pkg

import "fmt"

// сервис который сохраняет данные
type DataService struct {
	Next Service
}

func (upd DataService) Execute(data *Data) {
	if !data.UpdateSource {
		fmt.Printf("Data wasn't updated\n")
		return
	}
	fmt.Println("Data save")
}

func (upd *DataService) SetNext(svc Service) {
	upd.Next = svc
}
