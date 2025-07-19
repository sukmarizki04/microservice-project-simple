package novalagunggolang

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func mainExample() {
	var input string
	fmt.Print("Type some number :")
	fmt.Scanln(&input)
	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "isnumber")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())

	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Cannot be Empty")
	}
	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("error occured", r)
	} else {
		fmt.Println("Application running perfecly ")
	}
}

func mainFunction() {
	defer catch()

	var name string
	fmt.Printf("Type your name: ")
	fmt.Scanln(&name)
	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		fmt.Println(err.Error())
		panic(err.Error())
	}
}

// memanfaatkan error pada IIFE
func ExampleMain() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic Occured ", r)
		} else {
			fmt.Println("Application Running Perfecly")
		}
	}()
	panic("some error happen")
}

// perulangan panic dan error
func mainFungcion() {
	data := []string{"superman", "aquaman", "wonder woman"}
	for _, each := range data {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic occured on looping ", each, "| message:")
				} else {
					fmt.Println("Application running perfecly")
				}
			}()
			panic("some error happen")

		}()
	}
}
