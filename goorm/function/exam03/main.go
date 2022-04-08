package main

import "fmt"

const g float32 = 9.8

type Energy func(float32, float32) float32

func calMechEnergy(f Energy, a float32, b float32) (result float32) {
	result = f(a, b)
	return result
}

func main() {
	var m, v, h float32
	fmt.Scanln(&m, &v, &h)

	kinEnergy := func(m, v float32) float32 {
		return float32(1/2) * m * (v * v)
	}
	potEnergy := func(m, h float32) float32 {
		return m * g * h
	}

	ke := calMechEnergy(kinEnergy, m, v)
	pe := calMechEnergy(potEnergy, m, h)
	fmt.Printf("%f %f %f", ke, pe, ke+pe)
}
