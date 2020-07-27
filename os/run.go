package main

import (
	"github.com/reiver/cyber80/os/export"
	"github.com/reiver/cyber80/os/log"
	"github.com/reiver/cyber80/os/nextfunc"

	"github.com/reiver/go-c80"
	"github.com/containous/yaegi/interp"

	"fmt"
	"syscall/js"
)

// run() is run when the user click the ‘run’ button in the web browser.
//
// When that happens, we get the Go code that the user wrote in the
// textarea and try to interpret it (with a Go interpreter).
//
// If the interpretation is successful, it will look in the interpretted
// code for a “main.next” func. If it can find it, it will set “window._next”
// (in the web browser) to this it.
//
// The effect that will have is that that function will run about 60 times
// a second.
//
// The purpose of this is to update things to draw the next frame in the
// <canvas>.
//
// I.e., “main.next” will be the new animation loop function.
func run(this js.Value, args []js.Value) interface{} {
	log.Publish("run(): BEGIN")



	log.Publish("run(): creating interpreter")
	interpreter := interp.New(interp.Options{})
	if nil == interpreter {
		log.Publish("ERROR: Could not create interpreter.")
		panic("ERROR: Could not create interpreter.")
	}
	log.Publish("run(): interpreter created")



	log.Publish("run(): making packages available for import")
	interpreter.Use(export.Symbols)
	log.Publish("run(): packages made available for import")



	log.Publish("run(): getting code")
	src := textarea.Get("value").String()
	log.Publish("run(): code gotten")
	log.Publish("run(): code:", src)



	log.Publish("run(): interpreting code")
	_, err := interpreter.Eval(src)
	if nil != err {
		msg := fmt.Sprintf("ERROR: could not eval code: %s", err)

		log.Publish(msg)

		c80.Terminal.Publish(msg)
		nextfunc.Set(textnext)

		return nil
	}
	log.Publish("run(): code interpreted")



	log.Publish("run(): looking for main.next func")
	v, err := interpreter.Eval("main.next")
	if nil != err {
		msg := fmt.Sprintf("ERROR: could locate the main.next func: %s", err)

		log.Publish(msg)

		c80.Terminal.Publish(msg)
		nextfunc.Set(textnext)

		return nil
	}
	log.Publish("run(): main.next func found")



	log.Publish("run(): lifting main.next func")
	next, ok := v.Interface().(func())
	if !ok {
		msg := fmt.Sprintf("ERROR: main.next in not a func")

		log.Publish(msg)

		c80.Terminal.Publish(msg)
		nextfunc.Set(textnext)

		return nil
	}
	log.Publish("run(): main.next func lifted")



	log.Publish("run(): setting main.next func as new magic function")
	nextfunc.Set(next)
	log.Publish("run(): main.next func set as new magic function")



	log.Publish("run(): END")

	return nil
}
