package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// Run 'write' function synchonously
	i := 0
	for i < 10 {
		write("direct " + strconv.Itoa(i))
		i++
	}

	// Run 'write' function in a goroutine
	i = 0
	for i < 10 {
		go write("goroutine " + strconv.Itoa(i))
		i++
	}
	time.Sleep(50000000)
	fmt.Println("Finished")
}

func write(from string) {
	fmt.Println(from)
}
