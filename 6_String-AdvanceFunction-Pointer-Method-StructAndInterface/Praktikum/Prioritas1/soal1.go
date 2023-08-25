package main

import "fmt"

type Car struct {
	carType string
	fuelin  float64
}

func (c Car) jarakEstimasi() float64 {
	jarakTempuh := 1.5 // kilometer per milliliter (1.5 L/mill)
	jarakEstimasi := c.fuelin * jarakTempuh
	return jarakEstimasi
}

func main() {
	car := Car{
		carType: "SUV",
		fuelin:  10.5,
	}

	jarakEstimasi := car.jarakEstimasi()
	fmt.Printf("Car type: %s, est.distance: %.2f\n", car.carType, jarakEstimasi)
}
