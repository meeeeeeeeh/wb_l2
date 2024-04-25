package pkg

type Factory struct {
	Collector Collector
}

//передача интерфейса в качестве аргумента чтобы можно было создавать любую модель
func NewFactory(collector Collector) *Factory {
	return &Factory{Collector: collector}
}

func (factory *Factory) SetCollector(colector Collector) {
	factory.Collector = colector
}

//основная функция которая возвращает комплектованную сборку
//происходит пошагово
func (factory Factory) CreateComputer() Computer {
	factory.Collector.SetCore()
	factory.Collector.SetMemory()
	factory.Collector.SetBrand()
	factory.Collector.SetGraphicCard()
	factory.Collector.SetMonitor()
	return factory.Collector.GetComputer()
}