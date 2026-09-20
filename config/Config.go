package Config

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
)

type WindowApp struct {
	ColorWindow    color.Color
	Pathbackground string
	Pathbuycard    string
	NameWindow     string
	PosX           float32
	PosY           float32
	StartPosX      int
	StartPosY      int
}

type Buycard struct {
	Name  string
	Price int
	Photo *canvas.Image
	SizeX float32
	SizeY float32
	PosX  float32
	PosY  float32
}
