package main

import "fmt"

func main() {
	// var fullName string
	// var age int
	// var gpa float32
	// var country string

	// fullName = "Md Raihan Hossin"
	// age = 30
	// gpa = 3.47
	// country = "Bangladesh"

	// var fullName string = "Md Raihan Hossin";
	// var age int = 30;
	// var gpa float32 = 3.47;
	// var country string = "Bangladesh";

	fullName 	:= "Md Raihan Hossin";
	country  	:= "Bangladesh";
	age 		:= 30;
	gpa 		:= 3.47;

	fmt.Println(fullName, "is a student")
	fmt.Println(fullName, "is", age, "years old")
	fmt.Println(fullName, "has got", gpa, "/4 in bsc")
	fmt.Println(fullName, " is originally from", country)
}
