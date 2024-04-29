//адаптер позволяет разным объектам пакета использовать один функционал (интерфейс)

// все данные имеют формат json
// а чтобы работать с сервисом аналитики нужен xml
package data

import "fmt"

type JsonDocument struct {

}

func (doc JsonDocument) ConvertToXml() string {
	return "<xml>/<xml>"
}

type JsonDocumentAdapter struct {
	//добавляем все поля стрктуры
	JsonDocument *JsonDocument
}

func (adapter JsonDocumentAdapter) SendXmlData() {
	adapter.JsonDocument.ConvertToXml()
	fmt.Println("Sending xml data")
}