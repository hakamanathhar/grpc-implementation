package main

import (
	"fmt"
)

type CalculatorInterface interface {
	Add(a int, b int) int
}

type Calculator struct {
	filter FilterInterface
}

func NewCalculatorWithFilter(filter FilterInterface) CalculatorInterface {
	return &Calculator{
		filter: filter,
	}
}

func (c *Calculator) Add(a int, b int) int {
	// fmt.Printf("reflect Add c.calc: %v\n", reflect.TypeOf(c.iface))
	checked := c.filter.Check(a)
	fmt.Printf("checked: %v \n\n", checked)
	if checked {
		return a + b
	} else {
		return 0
	}
}
