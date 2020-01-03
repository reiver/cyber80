package main

import (
	"github.com/reiver/cyber80/src/cyber80/display"

	"fmt"
	"syscall/js"
)

func main() {

	fmt.Printf("Hello world! ⸻ %s\n", name)

	var window js.Value
	{
		window = js.Global()
	}

	var document js.Value
	{
		document = js.Global().Get("document")
	}

	var canvas js.Value
	{
		canvas = document.Call("createElement", "canvas")

		canvas.Call("setAttribute", "width",  cyber80_display.Width)
		canvas.Call("setAttribute", "height", cyber80_display.Height)

		canvas.Set("style", "position:absolute; left:0px; top:0px;")

		document.Get("body").Call("appendChild", canvas)
	}

	var ctx js.Value
	{
		ctx = canvas.Call("getContext", "2d")
	}

	var count int
	var redraw func()
	{
		redraw = func() {
			fmt.Printf("%s: redraw count: %d\n", name, count)

			ctx.Set("fillStyle", "#002b36") // solarized base03
			ctx.Call("fillRect", 0, 0, canvas.Get("width").Int(), canvas.Get("height").Int())

			msg := fmt.Sprintf("Hello world! ⸻ %d", count)
			count++

			ctx.Set("fillStyle", "#dc322f") // solarized red
			ctx.Call("fillText", msg, 10, 10)
		}
	}

	var resize func(this js.Value, args []js.Value)interface{}
	{
		resize = func(this js.Value, args []js.Value) interface{} {
			innerWidth  := window.Get("innerWidth").Int()
			innerHeight := window.Get("innerHeight").Int()

			newWidth := innerWidth
			newHeight := newWidth * cyber80_display.Height / cyber80_display.Width
			if newHeight > innerHeight {
				newWidth  = innerHeight * cyber80_display.Width / cyber80_display.Height
				newHeight = innerHeight
			}

			canvas.Set("width",  newWidth)
			canvas.Set("height", newHeight)
			fmt.Printf("%s: resize (width, height) = (%d, %d)\n", name, newWidth, newHeight)

			ctx.Call("scale", float64(newWidth) / float64(cyber80_display.Width), float64(newHeight) / float64(cyber80_display.Height))

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
		fmt.Printf("%s: hanging around...\n", name)
		ch := make(chan struct{})
		<-ch
		fmt.Printf("%s: Goodbye · Khodafez · 안녕 · 再見\n", name)
	}
}
