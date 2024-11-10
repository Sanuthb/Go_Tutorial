package main

import (
	"errors"
	"fmt"
)

func main() {
	printMatrix()
	var result1, result2, err = printMatrix2(5, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Division: %d, Remainder: %d\n", result1, result2)
	}
}

func printMatrix() {
	fmt.Println("hello world")
}

func printMatrix2(a int, b int) (int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("division by zero")
		return 0, 0, err
	}
	var division int = a / b
	var reminder int = a % b
	return division, reminder, err
}
