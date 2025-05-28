package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func printMessageAtRandomTime(id int, text string, channel chan string) {
	min := 0
	max := 3

	for i := range 10 {
		t := rand.Intn(max-min+1) + min
		log.Printf("[WORKER: %v] [%v] - I waited: %vs : %s", id, i, t, text)
		time.Sleep(time.Duration(t) * time.Second)
	}

	channel <- fmt.Sprintf("[WORKER: %v] now i'm done! bye!", id)

}

func main() {
	channel := make(chan string)
	numWorkers := 3

	for i := range numWorkers {
		go printMessageAtRandomTime(i, "Help meee!", channel)
	}

	for range numWorkers {
		msg := <-channel
		fmt.Println(">>> received:", msg)
	}

	fmt.Println(">>> All workers finished!")
}
