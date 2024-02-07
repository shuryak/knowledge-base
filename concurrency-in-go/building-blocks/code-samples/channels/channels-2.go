package main

func main() {
	writeStream := make(chan<- interface{})
	readStream := make(<-chan interface{})

	<-writeStream
	readStream <- struct{}{}
}
