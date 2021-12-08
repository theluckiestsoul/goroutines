package main

import (
	"fmt"
	"strconv"
	"time"
)

var messages chan string

func reader() {
	for {
		message := messages
		fmt.Println("Message received: " + <-message)
	}
}

func writer() {
	i := 0
	for i < 10 {
		m := "goroutine " + strconv.Itoa(i)
		messages <- m
		i++
	}
}

func main() {

	messages = make(chan string)
	go reader()
	go writer()

	time.Sleep(10 * time.Second)
	fmt.Println("Finished")
}
