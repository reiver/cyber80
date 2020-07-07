package main

import (
	"github.com/reiver/cyber80/src/cyber80/palette"
	"github.com/reiver/cyber80/src/cyber80/ram"
	"github.com/reiver/cyber80/src/cyber80/raster"

	"fmt"
	"syscall/js"
)

func main() {
	defer func() {
		if r := recover(); nil != r {
			log("CRASHED! 💀")
			logf("DEBUG: (%T) %#v", r, r)
		}
	}()

	log("Hello world!")
	logf("raster resolution: %d×%d", cyber80_raster.Width, cyber80_raster.Height)
	logf("RAM size: %d", cyber80_ram.Size)

	var window js.Value
	{
		window = js.Global()

		if js.Null() == window {
			log("ERROR: ‘window’ is null")
			return
		}
		if js.Undefined() == window {
			log("ERROR: ‘window’ is undefined")
			return
		}
	}

	var document js.Value
	{
		document = js.Global().Get("document")

		if js.Null() == document {
			log("ERROR: ‘document’ is null")
			return
		}
		if js.Undefined() == document {
			log("ERROR: ‘document’ is undefined")
			return
		}
	}

	var canvas js.Value
	{
		canvas = document.Call("createElement", "canvas")

		if js.Null() == canvas {
			log("ERROR: ‘canvas’ is null")
			return
		}
		if js.Undefined() == canvas {
			log("ERROR: ‘canvas’ is undefined")
			return
		}
	}

	var ctx js.Value
	{
		ctx = canvas.Call("getContext", "2d")

		if js.Null() == ctx {
			log("ERROR: ‘ctx’ is null")
			return
		}
		if js.Undefined() == ctx {
			log("ERROR: ‘ctx’ is undefined")
			return
		}
	}

	// Set up display.
	{
		canvas.Call("setAttribute", "width",  cyber80_raster.Width)
		canvas.Call("setAttribute", "height", cyber80_raster.Height)

		canvas.Set("style", "position:absolute; left:0px; top:0px;")

		document.Get("body").Call("appendChild", canvas)
	}

	// Splash
	splash()

	var redraw func()
	{
		redraw = func() {

			palette := cyber80_ram.Value.Palette()

			var colors [cyber80_palette.Count]string
			for i:=byte(0); i<cyber80_palette.Count; i++ {
				rgb := palette.RGB(i)
				colors[i] = fmt.Sprintf("#%x", rgb)
				//logf("palette colors[%d] = %q", i, colors[i])
			}

			{
				for i:=0; i<cyber80_raster.Size; i++ {
					value := cyber80_ram.Value[i]
					x, y := cyber80_raster.XY(i)

					ctx.Set("fillStyle", colors[value & 0x0f])
					ctx.Call("fillRect", x, y, 1, 1)
				}
			}
		}
	}

	var resize func(this js.Value, args []js.Value)interface{}
	{
		resize = func(this js.Value, args []js.Value) interface{} {
			innerWidth  := window.Get("innerWidth").Int()
			innerHeight := window.Get("innerHeight").Int()

			newWidth := innerWidth
			newHeight := newWidth * cyber80_raster.Height / cyber80_raster.Width
			if newHeight > innerHeight {
				newWidth  = innerHeight * cyber80_raster.Width / cyber80_raster.Height
				newHeight = innerHeight
			}

			logf("resize to (width, height) = (%d, %d)", newWidth, newHeight)

			canvas.Set("width",  newWidth)
			canvas.Set("height", newHeight)

			ctx.Call("scale", float64(newWidth) / float64(cyber80_raster.Width), float64(newHeight) / float64(cyber80_raster.Height))

			redraw()

			return nil
		}
	}
	resize(js.Null(), nil)

	{
		jsFunc := js.FuncOf(resize)

		window.Call("addEventListener", "resize", jsFunc, false)
	}

	{
		log("I am alive.")
		ch := make(chan struct{})
		<-ch
		log("Goodbye · Khodafez · 안녕 · 再見")
	}
}
