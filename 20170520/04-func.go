package main

import (
	"errors"
	"fmt"
)

func hello(a, b int) (int, bool, error) {
	if a > b {
		return 1, true, errors.New("Error: a > b")
	}

	return 0, false, errors.New("Error: a < b")
}

func main() {
	_, check, _ := hello(1, 2)
	if check {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	if _, _, err := hello(1, 2); err != nil {
		fmt.Println("true")
	} else {
		fmt.Println(err)
	}

	if _, _, err := hello(1, 2); err == nil {
		fmt.Println("true")
	} else {
		fmt.Println(err)
	}
}
