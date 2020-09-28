package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "A"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "B"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("OK: ", msg1)
		case msg2 := <-ch2:
			fmt.Println("OK: ", msg2)
		}
	}

	elapse := time.Since(start)
	fmt.Println(elapse)
}
