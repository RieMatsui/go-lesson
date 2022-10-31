package main

import (
	"fmt"
	"time"
)

func reciever(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	ch1 := make(chan int, 2)

	ch1 <- 1

	close(ch1)

	//ch1 <- 1
	//fmt.Println(<-ch1)

	i, ok := <-ch1
	fmt.Println(i, ok)

	i2, ok := <-ch1
	fmt.Println(i2, ok)

	ch2 := make(chan int, 2)

	go reciever("1.goroutine", ch2)
	go reciever("2.goroutine", ch2)
	go reciever("3.goroutine", ch2)

	y := 0
	for y < 100 {
		ch2 <- y
		y++
	}
	close(ch2)
	time.Sleep(3 * time.Second)
}
