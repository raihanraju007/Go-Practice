package main

import "fmt"

func main() {
	fmt.Println("Go methods practice")

	raihan := User{"Raihan","raihanraju007@gmail.com",true,30}
	// fmt.Println(raihan)
	fmt.Printf("Raihan details are: %+v\n",raihan)
	fmt.Printf("Name is: %v\n Email is: %v\n Status is: %v\n Age is: %v\n",raihan.Name,raihan.Email,raihan.Status,raihan.Age)

	raihan.getStatus()
	raihan.NewMail()

}

type User struct {
	Name 	string
	Email 	string
	Status 	bool
	Age 	int
}

func (u User) getStatus(){
	fmt.Println("Is user active: ", u.Status)
}

func(u User) NewMail(){
	u.Email = "raju@gmail.com"
	fmt.Println("Email of this user is: ", u.Email)
}
