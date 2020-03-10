package main

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

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

type Image struct {
	dx     int
	dy     int
	matrix [][]uint8
}

func NewImage(dx, dy int) Image {
	return Image{dx, dy, Pic(dx, dy)}
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.dx, i.dy)
}

func (i Image) At(x, y int) color.Color {
	value := i.matrix[y][x]
	return color.RGBA{value, value, 255, 255}
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	m := NewImage(100, 100)
	pic.ShowImage(m)
	fmt.Println(m.Bounds())
}
