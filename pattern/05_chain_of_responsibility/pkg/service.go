package pkg

type Service interface {
	Execute(*Data)
	SetNext(Service)
}

// содержит два флага:
// были ли переданы дынные - сервис принимающий данные
// савит сервис который обрабтал данные
type Data struct {
	GetSource    bool
	UpdateSource bool
}
