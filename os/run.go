package main

import (
	"github.com/reiver/cyber80/os/export"
	"github.com/reiver/cyber80/os/webbed"

	"github.com/containous/yaegi/interp"

	"syscall/js"
)

// run is run when the user click the ‘run’ button in the web browser.
//
// When that happens, run we get the Go code that the user wrote in the
// textarea and try to interpret it (with a Go interpreter).
//
// The the interpretation is successful, it will look in the interpretted
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
	log("run(): BEGIN")



	log("run(): creating interpreter")
	interpreter := interp.New(interp.Options{})
	if nil == interpreter {
		log("ERROR: Could not create interpreter.")
		panic("ERROR: Could not create interpreter.")
	}
	log("run(): interpreter created")



	log("run(): making packages available for import")
	interpreter.Use(export.Symbols)
	log("run(): packages made available for import")



	log("run(): getting code")
	src := textarea.Get("value").String()
	log("run(): code gotten")
	log("run(): code:", src)



	log("run(): interpreting code")
	_, err := interpreter.Eval(src)
	if nil != err {
		logf("ERROR: could not eval code: %s", err)
		panic(err)
	}
	log("run(): code interpreted")



	log("run(): looking for main.next func")
	v, err := interpreter.Eval("main.next")
	if nil != err {
		logf("ERROR: could eval for main.next: %s", err)
		panic(err)
	}
	log("run(): main.next func found")



	log("run(): lifting main.next func")
	next, ok := v.Interface().(func())
	if !ok {
		log("ERROR: could not life main.next func")
		panic("could not life main.next func")
	}
	log("run(): main.next func lifted")



	log("run(): setting main.next func as new magic function")
	{
		fn := func(this js.Value, args []js.Value) interface{} {
			next()

			return nil
		}

		jsFunc := js.FuncOf(fn)

		webbed.Window.Set(magicFunctionName, jsFunc)
	}
	log("run(): main.next func set as new magic function")



	log("run(): END")

	return nil
}
