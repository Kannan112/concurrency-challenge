package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- "abhinand"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "vimal"
	}()
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-ch1:
			fmt.Println(m1)
		case m2 := <-ch2:
			fmt.Println(m2)
		}
	}
}
