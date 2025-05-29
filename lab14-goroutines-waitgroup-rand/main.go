package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func printMessageAtRandomTimes(id int, text string, wg *sync.WaitGroup) {
	defer wg.Done()

	min := 0
	max := 3

	for i := range 10 {
		t := rand.Intn(max-min+1) + min
		log.Printf("[WORKER %v] [%v] [%v seconds]: %s", id, i, t, text)
		time.Sleep(time.Duration(t) * time.Second)
	}

}

func main() {

	worker := 3

	var wg sync.WaitGroup

	for i := range worker {
		wg.Add(1)
		go printMessageAtRandomTimes(i, "test message", &wg)
	}

	wg.Wait()

	fmt.Println("Finishing gracefuly...")

}
