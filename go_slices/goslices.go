package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the go slices practice program")

	var fruitList = [] string{"Apple", "Pineapple", "Orange"}
	// fmt.Printf("Type of fruit list is %T\n", fruitList)
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int,4)
	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867
	// highScore[4] = 123

	highScore = append(highScore, 555,666,777)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)

}