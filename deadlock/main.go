package main

import "fmt"

func main() {
	var ch chan string
	ch = make(chan string)
	ch <- "data"
	fmt.Println("ended")
}
