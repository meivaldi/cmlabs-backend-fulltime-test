package main

import "fmt"

func FizzBuzz(i int) (bool, string) {
	if (i%5 == 0) && (i%3 == 0) {
		return true, "FizzBuzz"
	} else if i%3 == 0 {
		return true, "Fizz"
	} else if i%5 == 0 {
		return true, "Buzz"
	}

	return false, ""
}

func main() {
	const n = 100

	var fizzBuzz bool
	var msg string
	for i := 1; i <= n; i++ {
		fizzBuzz, msg = FizzBuzz(i)
		if fizzBuzz {
			fmt.Println(msg)
			continue
		}

		fmt.Println(i)
	}
}
