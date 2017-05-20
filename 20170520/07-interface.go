package main

import (
	"fmt"

	"github.com/go-training/workshops/20170520/lexus"
	"github.com/go-training/workshops/20170520/toyota"
)

type cart interface {
	Total() float32
}

func Output(c cart) {
	fmt.Println(c.Total())
}

func main() {
	car1 := lexus.NewCart(1000)
	car2 := toyota.NewCart(2000)

	Output(car1)
	Output(car2)
}
