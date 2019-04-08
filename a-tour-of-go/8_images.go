package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
	//"math"
)

type Image struct {
	width, height  int
	colorA, colorB uint8
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i *Image) At(x, y int) color.Color {
	mutate := func(x, y int) uint8 {
		return uint8(x*x + y*y)
	}
	v := mutate(x, y)
	return color.RGBA{i.colorA, v, i.colorB, 255}
}

func main() {
	m := &Image{256, 256, 127, 255}
	pic.ShowImage(m)
}
