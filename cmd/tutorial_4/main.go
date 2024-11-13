package main

import "fmt"

func main() {
	switch 8 {
		case 8:
            fmt.Println("Number is between 0 and 5")
        case 6:
            fmt.Println("Number is between 6 and 10")
        default:
            fmt.Println("Number is outside the range")
	}

}