package main

import "fmt"

func sum(start int, end int, c chan int) {
	var sum = 0

	for i := start; i < end; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}

	c <- sum
}

func main() {
	c := make(chan int)

	go sum(0, 25, c)
	go sum(25, 50, c)

	x := <-c
	y := <-c

	fmt.Println(x, y, x+y)
}
