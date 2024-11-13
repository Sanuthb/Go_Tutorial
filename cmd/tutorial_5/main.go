package main

import "fmt"

func main() {
	var intarray [3]int32 //int32 is 4bytes 
	intarray[0] = 10
	intarray[1] = 20
	intarray[2] = 30
	fmt.Println(intarray)
	fmt.Println(intarray[1:3])

	fmt.Println(&intarray[1])// memory address 

	intarr := []int32{1,2,3,4}
	fmt.Println(intarr)

	var intslicearr []int32 = []int32{1,2,3}
	fmt.Printf("length of array: %v capacity of array %v",len(intslicearr),cap(intslicearr))
	intslicearr = append(intslicearr,4)
	fmt.Printf("\n length of array: %v capacity of array %v",len(intslicearr),cap(intslicearr))

	var intslicearr2 []int32 = []int32{4,5,6}
	intslicearr = append(intslicearr, intslicearr2...) // we can directly append intslicearr2 to the insslicearr 
	//error: cannot use intslicearr2 (variable of type []int32) as int32 value in argument to append

	fmt.Println(intslicearr)
	var insslicearr3 []int32 = make([]int32,2,4)
	insslicearr3 = append(intslicearr2,1,2,3,4,5)
	fmt.Println(insslicearr3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)
	var myMap2 = map[string]uint8{"adam":23,"jhon":40}
	myMap2["peter"] = 30
	fmt.Println(myMap2)
	fmt.Println(myMap2["peter"])
	var age,ok = myMap2["adam"]
	if ok {
		fmt.Println("age of adam is ",age)
	}else {
		fmt.Println("adam is not present in the map")
	}
	delete(myMap2,"adam")
	fmt.Println(myMap2)

	// for loop iteration
	// method 1
	for i := 0; i < len(intslicearr); i++ {
        fmt.Println(intslicearr[i])
    }
	// method 2
	for index,value := range intslicearr {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}
	// method 2
	i:=0
	for {
		if i>=len(intslicearr){
			break
		} 
		fmt.Println(intslicearr[i])
		i=i+1
	}

}

