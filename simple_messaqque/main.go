package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// using chanel
func worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Printf("Working area %d proccessing job %d\n ", id, job)
		time.Sleep(time.Second)
		result <- job * 2
	}
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbum(c *gin.Context) {
}
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbum)
	router.Run("localhost:8080")

}
