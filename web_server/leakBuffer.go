package main

var freeList = make(chan *Buffer, 100)
var serverchan = make(chan *Buffer)

func client() {
	for {
		var b *Buffer
		select {
		case b = <-freeList:
		default:
			b = new(Buffer)
		}
		load(b)
		serverchan <- b
	}
}
