package main

import (
	"fmt"
	// "sort"
)

func main() {
	// fmt.Println("Welcome to the go slices practice program")

	// var fruitList = [] string{"Apple", "Pineapple", "Orange"}
	// // fmt.Printf("Type of fruit list is %T\n", fruitList)
	// fruitList = append(fruitList, "Mango", "Banana")
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)

	// highScore := make([]int,4)
	// highScore[0] = 234
	// highScore[1] = 945
	// highScore[2] = 465
	// highScore[3] = 867
	// // highScore[4] = 123

	// highScore = append(highScore, 555,666,777)

	// fmt.Println(highScore)

	// sort.Ints(highScore)
	// fmt.Println(highScore)


	// creating slice
	// var mySlice [] int
	// fmt.Println(mySlice) 
	// fmt.Println(len(mySlice))
	// fmt.Printf("%T\n", mySlice)


	
	// slice from array
	// arr := [] int{10, 11, 12, 13, 14, 15}
	// mySlice := arr[1:4]
	// fmt.Println(mySlice) 
	// fmt.Println(len(mySlice))
	// fmt.Printf("%T\n", mySlice)


	// slice initialization
	//var mySlice []int = []int{10, 11, 12, 13, 14, 15} 
	// mySlice := []int{10, 11, 12, 13, 14, 15}
	// fmt.Println(mySlice) 
	// fmt.Println(len(mySlice))
	// fmt.Printf("%T\n", mySlice)

	// using make
	// mySlice := make([]int,3)
	// fmt.Println(mySlice)
	// fmt.Println(len(mySlice))
	// fmt.Printf("%T\n", mySlice)

	// mySlice := []int{1,2,3,4,5}

	//accessing slice
	// fmt.Println(mySlice[2])

	//modify slice
	// mySlice[2] = 60
	// fmt.Println(mySlice[2])

	// append element 
	// mySlice =  append(mySlice,100,200,300)
	// fmt.Println(mySlice)

	// Slice of slice
	// newSlice := mySlice[2:7]
	// fmt.Println(newSlice)

	// iterate over slice
	// for i, v := range mySlice {
	// 	fmt.Printf("Index: %d, Value: %d\n", i, v)
	// } 
	// copy a slice
	// mySlice := []int{1,2,3,4,5}
	// copyMySlice := mySlice
	// fmt.Println(mySlice)
	// fmt.Println(copyMySlice)

	// how to remove a value from slice based on index
	 var courses = []string{"reactjs", "javascript", "swift","python","ruby"}
	 fmt.Println(courses)

	 var index int = 2 
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)







}