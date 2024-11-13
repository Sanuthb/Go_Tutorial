package main

import "fmt"

 var globalvariable int = 10 // gloabal scope 

func main() {

	// placeholder used in printf 
	// 1. %d decimal point
	// 2. %f floating point
    // 3. %s string
    // 4. %c character
    // 5. %T type of the variable
    // 6. %b binary representation
    // 7. %x hexadecimal representation
    // 8. %o octal representation
	// 9. %p memory address
	// 10. %t boolean representation

    fmt.Printf("Hello, World!\n")
    fmt.Printf("Value of pi is %.2f\n", 3.14)
    fmt.Printf("Value of e is %.2e\n", 2.71828)
    fmt.Printf("Value of a is %d\n", 10)
    fmt.Printf("Value of b is %f\n", 20.5)
    fmt.Printf("Value of c is %c\n", 'A')

	var localvariable string = "a"
	fmt.Printf("Value of localvariable is %s\n", localvariable)
	fmt.Printf("Value of globalvariable is %d\n", globalvariable)
}