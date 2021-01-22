package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var i = os.Args[1]
	num, err := strconv.Atoi(i) // Convert i to integer

	if err != nil {
		fmt.Println(i, "is not a number")
		os.Exit(1)
	}

	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}
