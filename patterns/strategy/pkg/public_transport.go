package pkg

import "fmt"

type PublicTransportStrategy struct {
}

func (p *PublicTransportStrategy) Route(startPoint, endPoint int) {
	//параметры маршрута
	avgSpeed := 40
	total := endPoint - startPoint
	//40 - среднее время на маршрут
	totalTime := total * 40
	fmt.Printf("Public transport A: [%d] to B: [%d] Avg speed: [%d] Total: [%d] Total time: [%d] min\n", startPoint, endPoint, avgSpeed, total, totalTime)
}
