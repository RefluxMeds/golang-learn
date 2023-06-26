package main

import "fmt"

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int, 100)

	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)

	return c
}

func receive(src <-chan int) {
	for v := range src {
		fmt.Println(v)
	}
}
