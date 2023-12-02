package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func directCall(s string) {

	for i := 0; i < 3; i++ {
		fmt.Println("here", s)
		time.Sleep(1 * time.Millisecond)
	}

	wg.Done()

}

func main() {

	wg.Add(3)

	go directCall("abhi")

	go func(s string) {
		go directCall(s)
	}("nikil")

	FuncName := directCall

	go FuncName("joseph")
	wg.Wait()

}
