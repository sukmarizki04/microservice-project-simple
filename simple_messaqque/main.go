package main

import (
	"fmt"
	"time"
)

// using chanel
func worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Printf("Working area %d proccessing job %d\n ", id, job)
		time.Sleep(time.Second)
		result <- job * 2
	}
}
func main() {
	jobs := make(chan int, 5)
	result := make(chan int, 5)

	//start working
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, result)
	}

	//Send the jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	//Collecftion  result
	for a := 1; a <= 5; a++ {
		fmt.Printf("Result: %d\n", <-result)
	}

}
