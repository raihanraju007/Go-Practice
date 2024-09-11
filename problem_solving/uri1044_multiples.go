package main

import "fmt"

func main() {
	fmt.Println("beecrowd | 1044 Multiples")

	var A int
	var B int

	fmt.Scan(&A, &B) 
	if A % B == 0 || B % A == 0 {
		fmt.Println("Sao Multiplos")
		
	}else{
		fmt.Println("Nao sao Multiplos")
	}
}