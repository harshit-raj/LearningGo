package main

import (
	"fmt"
)

const usixteenbitmax float64 = 65535
const kmhMultiple float64 = 1.60934

type car struct {
	gasPedal      uint16
	brakePedal    uint16
	steeringWheel int16
	topSpeedKMH   float64
}

//by value
func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKMH / usixteenbitmax)

}

//by value
func (c car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKMH / usixteenbitmax / kmhMultiple)

}

//by pointer

func (c *car) newTopSpeed(newSpeed float64) {
	c.topSpeedKMH = newSpeed

}

func main() {
	aCar := car{gasPedal: 22341, brakePedal: 0, steeringWheel: 12561, topSpeedKMH: 225.0}
	// aCar := car{5000, 0, 12561, 225.0}
	fmt.Println(aCar.gasPedal)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
	aCar.newTopSpeed(500)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())

}
