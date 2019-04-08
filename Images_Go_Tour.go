package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)
type Image struct{
	w int
	h int 
}
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0,0,img.w, img.h)
}
func (img Image) At(x,y int) color.Color{
	v:=uint8(x*x+y*y)
	return color.RGBA{v,v,155,155}
}
func main() {
	m := Image{256,80}
	pic.ShowImage(m)
}
