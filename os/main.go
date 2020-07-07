package main

import (
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

	var window js.Value
	{
		window = js.Global()

		if window.IsNull() {
			log("ERROR: ‘window’ is null")
			return
		}
		if window.IsUndefined() {
			log("ERROR: ‘window’ is undefined")
			return
		}
	}

	var uint8Array js.Value
	{
		uint8Array = js.Global().Get("Uint8Array")

		if uint8Array.IsNull() {
			log("ERROR: ‘uint8Array’ is null")
			return
		}
		if uint8Array.IsUndefined() {
			log("ERROR: ‘uint8Array’ is undefined")
			return
		}
	}

	var document js.Value
	{
		document = js.Global().Get("document")

		if document.IsNull() {
			log("ERROR: ‘document’ is null")
			return
		}
		if document.IsUndefined() {
			log("ERROR: ‘document’ is undefined")
			return
		}
	}

//@TODO: Should this be hardcoded like this.
	var memory [256*256*4]byte

	var next func(this js.Value, args []js.Value) interface{}
	{
		next = func(this js.Value, args []js.Value) interface{} {

			for i:=0; i<len(memory); i++ {
				if 3 == i%4 {
					memory[i] = 255
				} else {
					memory[i] = byte(randomness.Intn(256))
				}
			}

			u8array := uint8Array.New(args[0])
			_ = js.CopyBytesToJS(u8array, memory[:])

			return nil
		}
	}

	{
		jsFunc := js.FuncOf(next)

		window.Set(magicFunctionName, jsFunc)
	}

	{
		log("I am alive!!")
		ch := make(chan struct{})
		<-ch
		log("Goodbye · Khodafez · 안녕 · 再見")
	}
}

