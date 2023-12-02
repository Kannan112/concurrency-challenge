package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		// In this case all of them print 4 since loop reaches 4 before executing the goroutine
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
		time.Sleep(5 * time.Second) // added time.sleep

	}

	wg.Wait()
	fmt.Println("finished")
}
