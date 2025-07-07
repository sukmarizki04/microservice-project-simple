package main

import (
	"fmt"
	"time"
)

var messageQue = make(chan string, 5)

func producer() {
	for i := 1; i <= 5; i++ {
		msg := fmt.Sprintf("Pesan ke-%d", i)
		fmt.Println("Producer: mengirim", msg)
		messageQue <- msg
		time.Sleep(time.Second)
	}
}
func consumer(id int) {
	for msg := range messageQue {
		fmt.Printf("Consumer %d: memproses %s \n", id, msg)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	go producer()
	go consumer(1)
	go consumer(2)

	time.Sleep(10 * time.Second)
}
