package pkg

import "fmt"

// устройство - источник данных

type Device struct {
	Name string
	Next Service
}

func (device *Device) Execute(data *Data) {
	//если данные уже были обработаны то вывоится сообщение об этом
	//если это было сделано то данные передаются на обработку другому сервису
	if data.GetSource {
		fmt.Printf("Data from device [%s] already get\n", device.Name)
		device.Next.Execute(data)
		return
	}
	fmt.Printf("Get data from device [%s]\n", device.Name)
	data.GetSource = true
	device.Next.Execute(data)
}

func (device *Device) SetNext(svc Service) {
	device.Next = svc
}
