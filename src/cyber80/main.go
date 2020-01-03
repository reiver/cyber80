package main

import (
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

		canvas.Call("setAttribute", "width",  displayWidth)
		canvas.Call("setAttribute", "height", displayHeight)

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
			newHeight := newWidth * displayHeight / displayWidth
			if newHeight > innerHeight {
				newWidth  = innerHeight * displayWidth / displayHeight
				newHeight = innerHeight
			}

			canvas.Set("width",  newWidth)
			canvas.Set("height", newHeight)
			fmt.Printf("%s: resize (width, height) = (%d, %d)\n", name, newWidth, newHeight)

			ctx.Call("scale", newWidth / displayWidth, newHeight / displayHeight)

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
