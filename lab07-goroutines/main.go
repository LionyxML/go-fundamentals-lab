package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println(text)
		time.Sleep(400 * time.Millisecond)
	}
	channel <- "Done!"
}

func main() { // main goroutine
	var channel chan string

	go printMessage("Jira sucks!", channel)

	response := <-channel

	fmt.Println(response)

}
