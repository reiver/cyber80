package main

import (
	"github.com/reiver/cyber80/os/webbed"

	"github.com/reiver/go-c80"

	"syscall/js"
)

const (
	magicFunctionName = "_next"
)

func main() {

	// This code will capture any panic().
	//
	// That way if this program crashes while running in the web browser
	// then we will get a nicer error message in the web browser's console.
	defer func() {
		if r := recover(); nil != r {
			log("CRASHED! 💀")
			logf("DEBUG: (%T) %#v", r, r)
		}
	}()

	// We output this message — in the web browser's console — so that
	// we provide some indication that this program has started successfully.
	log("Hello world!")


	var next func(this js.Value, args []js.Value) interface{} = nextfunc(func(){
		c80.Draw(c80.Terminal)
	})

	// Here we are assigning this “next” function in this Go program to:
	//
	//	window._next
	//
	// … in the JavaScript in the web browser.
	//
	// Remember that “window._next” is a sort of ‘magic’ location, in that
	// other parts of the program will run that every frame of the animation.
	{
		jsFunc := js.FuncOf(next)

		webbed.Window.Set(magicFunctionName, jsFunc)
	}

	// We need to prevent this program from ending.
	//
	// Here we do that by making the program wait (forever) on a
	// channel. This is a Go-trick. If you have never seen it
	// before you might think it looks a bit weird. But the purpose
	// of it is just to make it so this program does not end.
	{
		log("I am alive!!")
		ch := make(chan struct{})
		<-ch
		log("Goodbye · Khodafez · 안녕 · 再見")
	}
}
