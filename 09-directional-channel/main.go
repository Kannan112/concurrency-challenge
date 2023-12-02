package main

import "fmt"

func send(ch chan string) {
	ch <- "Hai"
}

func replay(ch1 chan string, ch2 chan string) {
	value := <-ch1
	ch2 <- value
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go send(ch1)
	go replay(ch1, ch2)

	fmt.Println(<-ch2)

}
