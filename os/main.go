package main

import (
	"github.com/reiver/cyber80/os/log"
	"github.com/reiver/cyber80/os/nextfunc"
)

func main() {

	// This code will capture any panic().
	//
	// That way if this program crashes while running in the web browser
	// then we will get a nicer error message in the web browser's console.
	defer func() {
		if r := recover(); nil != r {
			log.Publish("CRASHED! 💀")
			log.Publishf("DEBUG: (%T) %#v", r, r)
		}
	}()

	// We output this message — in the web browser's console — so that
	// we provide some indication that this program has started successfully.
	log.Publish("Hello world!")

	// Here we are assigning this “next” function in this Go program to:
	//
	//	window._next
	//
	// … in the JavaScript in the web browser.
	//
	// Remember that “window._next” is a sort of ‘magic’ location, in that
	// other parts of the program will run that every frame of the animation.
	nextfunc.Set(textnext)

	// We need to prevent this program from ending.
	//
	// Here we do that by making the program wait (forever) on a
	// channel. This is a Go-trick. If you have never seen it
	// before you might think it looks a bit weird. But the purpose
	// of it is just to make it so this program does not end.
	{
		log.Publish("I am alive!!")
		ch := make(chan struct{})
		<-ch
		log.Publish("Goodbye · Khodafez · 안녕 · 再見")
	}
}
