package main

import "fmt"

func main() {
	defer fmt.Println("Halo")
	fmt.Println(" Selamat Pagi ")

}

func orderSomeFood(menu string) {
	defer fmt.Println("Terimakasih telah menunggu, Silahkan Tunggu")
	if menu == pizza {
		fmt.Println("Pilihan tepat ", " ")
		fmt.Println("Pizza ditempat kami paling enak!", "\n")
		return
	}

	fmt.Println("pesanan anda:", menu)
}

func exampleMain() {
	number := 3
	if number == 3 {
		fmt.Println("halo 1")
		func() {
			defer fmt.Println("halo3")
		}()

	}
	fmt.Println("halo2")
}
