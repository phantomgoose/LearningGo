package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prevTwo := [2]int{0, 1}
	return func() int {
		defer func() {
			next := prevTwo[0] + prevTwo[1]
			prevTwo[0] = prevTwo[1]
			prevTwo[1] = next
		}()
		return prevTwo[0]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
