package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	raihan := User{"Raihan","raihanraju007@gmail.com",true,30}
	// fmt.Println(raihan)
	fmt.Printf("Raihan details are: %+v\n",raihan)
	fmt.Printf("Name is: %v\n Email is: %v\n Status is: %v\n Age is: %v\n",raihan.Name,raihan.Email,raihan.Status,raihan.Age)

}

type User struct {
	Name 	string
	Email 	string
	Status 	bool
	Age 	int
}