package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// init 2d slices
	matrix := make([][]uint8, dy)
	for i := range matrix {
		matrix[i] = make([]uint8, dx)
	}

	// populate matrix
	for y := range matrix {
		for x := range matrix[y] {
			matrix[y][x] = uint8(x * y)
		}
	}
	return matrix
}

func main() {
	pic.Show(Pic)
}
