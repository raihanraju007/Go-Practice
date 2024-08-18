package main

import "fmt"

func main() {
	// fmt.Println("Hello World From Go")
	var fullName string;
	var age int;
	var gpa float32;
	var country string;

	fullName = "Md Raihan Hossin";
	age = 30;
	gpa = 3.47;
	country = "Bangladesh";

	fmt.Println(fullName, "is a student");
	fmt.Println(fullName, "is", age, "years old");
	fmt.Println(fullName, "has got", gpa,"/4 in bsc");
	fmt.Println(fullName, " is originally from", country);

}
