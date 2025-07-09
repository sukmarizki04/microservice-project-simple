package main

import "fmt"

func main() {
	var secret interface{}

	secret = "Sukma Rizki"
	fmt.Println(secret)

	secret = map[string]int{"sukma": 1, "rizki": 2}
	fmt.Println(secret)
	var data map[string]interface{}

	data = map[string]interface{}{
		"name":      "Sukma",
		"age":       23,
		"breakFast": []string{"apple", "banana", "watermelon"},
	}
	fmt.Println(data)

}
