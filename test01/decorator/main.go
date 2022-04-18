package main

import "fmt"

type pizza interface {
	getPrice() int
}

type combinationPizza struct{}

func (c *combinationPizza) getPrice() int {
	return 30
}

type pineappleTopping struct {
	pizza pizza
}

func (p *pineappleTopping) getPrice() int {
	pizzaPrice := p.pizza.getPrice()
	return pizzaPrice + 10
}

type oliveTopping struct {
	pizza pizza
}

func (o *oliveTopping) getPrice() int {
	pizzaPrice := o.pizza.getPrice()
	return pizzaPrice + 5
}

func main() {
	combination := combinationPizza{}

	addPineapple := pineappleTopping{&combination}
	addOlive := oliveTopping{&addPineapple}

	fmt.Println(addOlive.getPrice())
}
