package novalagunggolang

import (
	"fmt"
	"runtime"
	"time"
)

func bufferedChanel() {
	runtime.GOMAXPROCS(2)
	messages := make(chan int, 3)
	go func() {
		for {
			i := <-messages
			fmt.Println("Menerima data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("Kirim data", i)
		messages <- i
	}
	time.Sleep(1 * time.Second)
}

//Chanel select
/*Chanel memebuat manajemen goroutine menjadi sangat mudah Namun perlu diingat , fungsi
utama chanel adalah bukan untuk kontrol eksekusi goroutine melainkan untuk sharing data
atau komunikasi goroutine atau pesan ,Tergantung jenis kasusnya, ada kalanya kita lebih dari stu chanel untuk komunikasi data antar goruitne
Penerimaaan data pada banyak goroutine penerapanya masih sama */
