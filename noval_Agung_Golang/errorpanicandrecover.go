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
