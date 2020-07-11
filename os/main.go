package main

import (
	"github.com/reiver/cyber80/os/webbed"

	"github.com/reiver/go-font8x8"

	"image"
	"image/draw"
	"syscall/js"
)

const (
	magicFunctionName = "_next"
)

func main() {
	defer func() {
		if r := recover(); nil != r {
			log("CRASHED! 💀")
			logf("DEBUG: (%T) %#v", r, r)
		}
	}()

	log("Hello world!")



	var next func(this js.Value, args []js.Value) interface{}
	{
		next = func(this js.Value, args []js.Value) interface{} {

			for offset, character := range text {
				cx := offset % 32
				cy := offset / 32

				x := cx*8
				y := cy*8

				rect := image.Rectangle{
					Min: image.Point{
						X:x,
						Y:y,
					},
//@TODO: Should these be hardcoded.
					Max: image.Point{
						X:x+8,
						Y:y+8,
					},
				}

				var src image.Image = font8x8.Rune(character)

//				draw.Draw(&frameImage, rect, &image.Uniform{color.RGBA{0, 0, 255, 255}}, image.ZP, draw.Src)
				draw.Draw(&frameImage, rect, src, image.ZP, draw.Src)
			}


//			for i:=0; i<len(memory); i++ {
//				if 3 == i%4 {
//					memory[i] = 255
//				} else {
//					memory[i] = byte(randomness.Intn(256))
//				}
//			}

			u8array := uint8Array.New(args[0])
			_ = js.CopyBytesToJS(u8array, frame[:])

			return nil
		}
	}

	{
		jsFunc := js.FuncOf(next)

		webbed.Window.Set(magicFunctionName, jsFunc)
	}

	{
		log("I am alive!!")
		ch := make(chan struct{})
		<-ch
		log("Goodbye · Khodafez · 안녕 · 再見")
	}
}
