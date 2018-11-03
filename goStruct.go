package main

import (
	"fmt"
)

type car struct {
	gasPedal      uint16
	brakePedal    uint16
	steeringWheel int16
	topSpeedKMH   float64
}

func main() {
	a_car := car{gasPedal: 5000, brakePedal: 0, steeringWheel: 12561, topSpeedKMH: 225.0}
	// a_car := car{5000, 0, 12561, 225.0}
	fmt.Println(a_car.gasPedal)

}
