package main

import "fmt"

func main() {
	fmt.Println("Starting array practice")

	var fruitList [4] string;
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Pineapple"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [3] string{"potato", "beans", "mushroom"}
	fmt.Println("Vegetable list is: ", vegList)
}