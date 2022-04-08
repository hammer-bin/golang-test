package main

import "fmt"

const g float32 = 9.8

type Energy func(float32, float32) float32

func calMechEnergy(f Energy, a float32, b float32) (result float32) {
	result = f(a, b)
	return
}

func main() {
	var m, v, h float32
	fmt.Scanln(&m, &v, &h)

	kinEnergy := func(m, v float32) float32 {
		return (m * v * v) / 2
	}
	potEnergy := func(m, h float32) float32 {
		return m * g * h
	}

	ke := calMechEnergy(kinEnergy, m, v)
	pe := calMechEnergy(potEnergy, m, h)
	fmt.Printf("%v %v %v", ke, pe, ke+pe)
}
