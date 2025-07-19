package novalagunggolang

import (
	"fmt"
	"math/rand"
)

func random() {
	//Pada chapter ini kita akan belajar pemanfaatan package math/rand untuk pembuatan data acak atau random
	ranomizer := rand.New(rand.NewSource(10))
	fmt.Println("random ke 1:", ranomizer.Int())
	fmt.Println("random ke 2:", ranomizer.Int())
	fmt.Println("random ke 3:", ranomizer.Int())

	/*S*/
}
