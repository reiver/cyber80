package main

import (
	"github.com/reiver/cyber80/os/webbed"

	"syscall/js"
)

const (
	// This should be the ‘id’ on the <button> element in the HTML code.
	buttonMagicID = "magic256-run"
)

var (
	button js.Value
)

func init() {
	button = webbed.Document.Call("getElementById", buttonMagicID)

	if button.IsNull() {
		msg := "ERROR: ‘button’ is null"

		log(msg)
		panic(msg)
	}
	if button.IsUndefined() {
		msg := "ERROR: ‘button’ is undefined"

		log(msg)
		panic(msg)
	}

	button.Call("addEventListener", "click", js.FuncOf(run))
}
