package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	toGet(&wg)
	wg.Wait()
	fmt.Println("Finished")
}

func toGet(wg *sync.WaitGroup) {

	var data int
	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Second)
		data++
		fmt.Println("value of the data", data)
	}()
	fmt.Println("return at this point")
	return
}
