package main

import (
	"fmt"
)

// go:generate goautomock -template=testify Filter
type FilterInterface interface {
	Check(a int) bool
}

type Filter struct{}

func (fl *Filter) Check(a int) bool {
	fmt.Printf("filter default\n")
	if a > 10 {
		return false
	}
	return true
}
