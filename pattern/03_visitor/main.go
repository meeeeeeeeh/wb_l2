/*
поведенческий паттерн, который позволяет добавить новую операцию
для целой иерархии классов, не изменяя код этих классов

пример:
в библиотеке есть структуры квадрат, круг, треугольник
нужно сделать так чтобы код можно было дополнять, минимально изменяя структуры
*/

package main

import (
	"fmt"
	"visitor/pkg"
)

func main() {
	square := &pkg.Square{Side: 2}
	circle := &pkg.Circle{Radius: 3}
	rectangle := &pkg.Rectangle{L: 2, B: 3}

	areaCalculator := &pkg.AreaCalculator{}

	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &pkg.MiddleCoordinates{}
	square.Accept(middleCoordinates)
	circle.Accept(middleCoordinates)
	rectangle.Accept(middleCoordinates)
}
