package pkg

import "fmt"

// сервис который обновляет данные от устройства
type UpdateDataService struct {
	Name string
	Next Service
}

func (upd *UpdateDataService) Execute(data *Data) {
	//если данные уже были обработаны то вывоится сообщение об этом
	//если это было сделано то данные передаются на обработку другому сервису
	if data.UpdateSource {
		fmt.Printf("Data inservice [%s] is already updated\n", upd.Name)
		upd.Next.Execute(data)
		return
	}
	fmt.Printf("Update data from service [%s]\n", upd.Name)
	data.UpdateSource = true
	upd.Next.Execute(data)
}

func (upd *UpdateDataService) SetNext(svc Service) {
	upd.Next = svc
}
