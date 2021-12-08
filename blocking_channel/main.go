package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan string
	ch = make(chan string)

	go func() {
		time.Sleep(20 * time.Second)
		fmt.Println("before read")
		fmt.Println(time.Now().UTC())

		<-ch

		fmt.Println("after read")
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("before write")
	fmt.Println(time.Now().UTC())

	ch <- "data"

	fmt.Println("after write")
}
