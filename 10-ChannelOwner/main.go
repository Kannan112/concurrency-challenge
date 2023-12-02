package main

import "fmt"

func main() {
	producer := func() <-chan int {
		ch := make(chan int)

		go func() {
			for i := 0; i < 4; i++ {
				ch <- i
			}
		}()

		return ch
	}

	consumer := func(ch <-chan int) {
		for r := range ch {
			fmt.Println("value", r)
		}

	}

	got := producer()
	consumer(got)
	close(ch)

}
