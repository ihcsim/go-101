package main

import "golang.org/x/tour/pic"

func main() {
	pic.Show(Pic)
}

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dx)

	for index := range picture {
		picture[index] = make([]uint8, dy)
	}

	for x := 0; x < len(picture); x++ {
		for y := 0; y < len(picture[x]); y++ {
			picture[x][y] = uint8(x * y)
		}
	}

	return picture
}
