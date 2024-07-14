package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Pyhton"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	// loops are intersting in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
