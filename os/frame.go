package main

import (
	"image"
)

const (
	frameWidth  = 256
	frameHeight = 256
	frameDepth  = 4   // RGBA
)

var frame [frameWidth*frameHeight*frameDepth]byte

var frameImage image.NRGBA = image.NRGBA{
	Pix: frame[:],
	Stride:frameWidth*frameDepth,
	Rect: image.Rectangle{
		Min: image.Point{
			X:0,
			Y:0,
		},
		Max: image.Point{
			X:(frameWidth-1),
			Y:(frameHeight-1),
		},
	},
}
