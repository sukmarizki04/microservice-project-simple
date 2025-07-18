package novalagunggolang

import (
	"fmt"
	"os"
)

func exitfunction() {
	defer fmt.Println("halo")
	os.Exit(1)
	fmt.Println("selamat datang")
}
