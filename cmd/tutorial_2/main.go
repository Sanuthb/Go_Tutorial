package main

import ("fmt")

func main() {
	//data types 
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// float32 float64
	// string
	// bool

	var a, b int = 10, 20
	sum := a + b
	fmt.Printf("Sum of %d and %d is %d\n", a, b, sum)

	var c =10 
	fmt.Printf("Value of c is %d\n", c)
	d:= 10 
    fmt.Println("Value of d:", d)
	
	const pi = 3.14
	fmt.Printf("Value of pi is %.2f\n", pi)
}