//поведенческий паттерн проектирования
//определяет схожие алгоритмы и и помещает каждый из них в свою отдельную структуру
//после чего алгоритмы могут взаимодействовать

//проблемы коорые решает:
//есть приложение навигатор для построения маршрутов
//нужно построить пеший маршрут, для велосипедов, на машине итд
//-> возникает много зависемостей от того какой маршрут выберет пользователя
//кроме того должен быть выбор маршрута и указать какой будет самым оптимальным на данном участке
//паттерн стратегия предлагает определить похожие алгоритмы которые часто
//изменяются или расширяются и вынести их в собственный объект и называть его стратегия
//построение маршрута - стратегия и каждая стратегия имеет свое поведение
//road, public transport, walking strategy

//паттерн нужно применять когда нужно использовать разные варианты
//алгоритма внутри одного объекта (навигатор)
//когда есть много похожих объектов отличающихся только определенным поведением
//когда не хочется раскрывать детали реализации общих алгоритмов
//когда есть различные варианты алгоритмов реализованные в виде условного оператора

//плюсы:
//замена алгоритмов налету
//изолирует код алгоритмов от остальной логики
//уход от наследования
//реализует принцип открытости и закрытости

//минусы:
//усложнение программы засчет допонительных объектов
//клиент должен знать в чем разница между стратегиями чтобы выбрать подходящую

package main

import "strategy/pkg"

var (
	start      = 10
	end        = 100
	strategies = []pkg.Strategy{
		&pkg.PublicTransportStrategy{},
		&pkg.RoadStrategy{},
		&pkg.WalkStrategy{},
	}
)

// получаем 3 варианта построения одного маршрута
func main() {
	nvg := pkg.Navigator{}
	for _, strategy := range strategies {
		nvg.SetStrategy(strategy)
		nvg.Route(start, end)
	}
}
