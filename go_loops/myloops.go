package main

import "fmt"

func main() {
	fmt.Println("Go loops practice")

	// days := [] string{"Sunday", "Tuesday","Wednesday","Friday","Saturday"}

	// fmt.Println(days)

	// for i:=0; i< len(days); i++{
	// 	fmt.Println(days[i])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }
	// for index, day := range days{
	// 	fmt.Printf("Index is %v and value is %v\n",index,day)
	// }

	rougueValue := 1
	for rougueValue < 10{

		for rougueValue ==2 {
			goto raju
		}

		if rougueValue == 5{
			break
		}
		fmt.Println("Value is: ",rougueValue)
		rougueValue++
	}

	raju:
		fmt.Println("Hello Raju!")

	
}