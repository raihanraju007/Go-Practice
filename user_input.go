package main

import "fmt"

func main(){

	var fullName string;
	var number1, number2 int;

	// Input : Scan, Scanln, Scanfageage
	// Output : Print, Println, Printf

	fmt.Print("Enter your name: ")
	fmt.Scan(&fullName);
	

	fmt.Print("Enter Two Numbers: ");
	fmt.Scan(&number1, &number2);
	// fmt.Scanf("%v %v", &number1, &number2);

	fmt.Printf("Full name is :  %v\n",fullName);
	fmt.Printf("number1 = %v, number2 = %v\n",number1,number2);


}