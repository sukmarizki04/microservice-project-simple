package main

import (
	"fmt"
	"reflect"
)

func combinationSliceMapAndInterface() {
	var person = []map[string]interface{}{
		{"name": "wick", "age": 23},
		{"name": "Budi", "age": 23},
		{"name": "jodi", "age": 23},
		{"name": "fira", "age": 23},
	}
	for _, each := range person {
		fmt.Println(each["name"], "age is", each["each"])
	}

	var fruits = []interface{}{
		map[string]interface{}{
			"name": "Straberry", "total": 10,
		}, []string{"manggo", "pineaple", "papaya"}, "orange",
	}

	for _, each := range fruits {
		fmt.Println(each)
	}

	number := 20
	var reflectionValue = reflect.ValueOf(number)
	fmt.Println("type of variabel: ", reflectionValue.Type())
	if reflectionValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel : ", reflectionValue.Int())
	}

	//Akses nilai dalam bentuk interface
	var angka = 23
	var reflectValue = reflect.ValueOf(angka)
	fmt.Println("Tipe Variabel :", reflectValue.Type())
	fmt.Println("nilai variabel :", reflectValue.Interface())
}

func main() {

}
