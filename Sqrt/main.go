package main

import (
	"fmt"
	"math"
)

func sqrt(x float64, accuracy float64) float64 {
	res, prev := 1.0, 0.0
	for {
		if prev != 0.0 && math.Abs(res/prev-1) < accuracy {
			break
		}
		prev = res
		res -= (res*res - x) / (2 * res)
	}
	return res
}

func main() {
	fmt.Println(sqrt(2, 0.00000001))
	fmt.Println(math.Sqrt(2))
}
