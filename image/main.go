package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

func main() {
	m := Image{}
	pic.ShowImage(m)
}

type Image struct{}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 150)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{150, 150, 255, 255}
}
