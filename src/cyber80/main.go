package main

import (
	"fmt"
	"syscall/js"
)

func main() {

	fmt.Println("Hello world! ⸻ cyber80")

	document := js.Global().Get("document")
	canvas := document.Call("createElement", "canvas")
	canvas.Call("setAttribute", "width",  displayWidth)
	canvas.Call("setAttribute", "height", displayHeight)
	document.Get("body").Call("appendChild", canvas)
}
