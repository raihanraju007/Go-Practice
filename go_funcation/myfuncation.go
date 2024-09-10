package main

import "fmt"

func main() {
	fmt.Println("Go function practice")
	greeter()

	result := adder(3,2)
	fmt.Println("Result is: ", result)

	proRes := proAdder(5,10,50,65,100,500)
	fmt.Println("Pro result is: ", proRes)

	
}


func greeter(){
	fmt.Println("Hello from golang")
} 

func adder(valOne int, valTwo int) int{
	return valOne + valTwo
}

func proAdder(values ...int) int{
	total := 0

	for _, val := range values{
		total += val
	}
	return total
}

