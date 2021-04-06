package main

import (
	"fmt"
	"strconv"
)

func main() {
	var out string
	var end = 100
	for i := 1; i <= end; i++ {
		if i%15 == 0 {
			out = "FizzBuzz"
		} else if i%3 == 0 {
			out = "Fizz"
		} else if i%5 == 0 {
			out = "Buzz"
		} else {
			out = strconv.Itoa(i)
		}
		if i != end {
			fmt.Printf("%s, ", out)
		} else {
			fmt.Print(out)
		}
	}
}
