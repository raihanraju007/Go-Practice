package main

import "fmt"

func main() { 
	
	defer fmt.Println("World")
	 defer fmt.Println("One")
	 defer fmt.Println("Two")

	fmt.Println("Hello")

}