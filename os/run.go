package main

import (
	"github.com/reiver/cyber80/os/export"
	"github.com/reiver/cyber80/os/webbed"

	"github.com/reiver/go-c80"
	"github.com/containous/yaegi/interp"

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
		return nil
	}
	log("run(): code interpreted")



	log("run(): looking for main.next func")
	v, err := interpreter.Eval("main.next")
	if nil != err {
		logf("ERROR: could locate the main.next func: %s", err)
		return nil
	}
	log("run(): main.next func found")



	log("run(): lifting main.next func")
	next, ok := v.Interface().(func())
	if !ok {
		log("ERROR: main.next in not a func")
		return nil
	}
	log("run(): main.next func lifted")



	log("run(): setting main.next func as new magic function")
	{
		fn := func(this js.Value, args []js.Value) interface{} {
			next()

			u8array := uint8Array.New(args[0])
			_ = js.CopyBytesToJS(u8array, c80.Frame())

			return nil
		}

		jsFunc := js.FuncOf(fn)

		webbed.Window.Set(magicFunctionName, jsFunc)
	}
	log("run(): main.next func set as new magic function")



	log("run(): END")

	return nil
}
