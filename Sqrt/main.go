package main

import (
	"fmt"
	"math"
	"strconv"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", strconv.FormatFloat(float64(e), 'f', 6, 64))
}

func sqrt(x float64, accuracy float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	res, prev := 1.0, 0.0
	for {
		if prev != 0.0 && math.Abs(res/prev-1) < accuracy {
			break
		}
		prev = res
		res -= (res*res - x) / (2 * res)
	}
	return res, nil
}

func main() {
	fmt.Println(sqrt(2, 0.00000001))
	fmt.Println(math.Sqrt(2))
	fmt.Println(sqrt(-1.1, 0.001))
}
