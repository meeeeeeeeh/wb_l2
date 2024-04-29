// внешний сервис который работает с форматом xml

package data_service

import "fmt"

type AalytialDataService interface {
	SendXmlData() interface{}
}

type XmlDoument struct {

}

func (doc XmlDoument) SendXmlData() {
	fmt.Println("Sending xml document")
}