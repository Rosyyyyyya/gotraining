package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int, 10)
	full_done := make(chan bool)
	defer close(ch)

	// producer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-full_done:
				time.Sleep(5 * time.Second)
				fmt.Printf("subprocess exit!\n")
				return
			default:
				n := rand.Intn(10)
				fmt.Println("generate element: ", n)
				ch <- n
				if len(ch) == 10 {
					fmt.Printf("current length of msg channel: %d, will exit subprocess\n", len(ch))
					close(full_done)
				}
			}
		}
	}()

	// consumer
	fmt.Println("begin consumer after 15s")
	time.Sleep(15 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		fmt.Printf("receive element from producer: %d\n", <-ch)
		if len(ch) == 0 {
			time.Sleep(5 * time.Second)
			fmt.Printf("current length of msg channel: %d, will break\n", len(ch))
			break
		}
	}
	fmt.Printf("final length of message channel: %d\n", len(ch))
	fmt.Println("main procee exit!")
}
