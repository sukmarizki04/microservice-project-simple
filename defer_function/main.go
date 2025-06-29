package main

import "fmt"

func main() {
	name := "Sukma Rizki"
	fmt.Println(name)
	for i := 0; i <= 10; i++ {
		fmt.Println("angka", i)
	}
	if name != "Sukma Rizki" {
		fmt.Println("Salah")
	} else {
		fmt.Println("Benar")
	}
}
