package main

import "fmt"

func main() {
	fmt.Println("Starting go maps examples...")

	language := make(map[string]string)

	language["JS"] = "Javascript"
	language["RB"] = "Ruby"
	language["PY"] = "Python"

	fmt.Println(language)

	for _, value := range language{
		fmt.Printf("For key v, value is %v\n", value)
	}

}