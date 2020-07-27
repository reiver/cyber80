package main

import (
	"github.com/reiver/cyber80/os/log"
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
		log.Panic("ERROR: ‘button’ is null")
		return
	}
	if button.IsUndefined() {
		log.Panic("ERROR: ‘button’ is undefined")
		return
	}

	button.Call("addEventListener", "click", js.FuncOf(run))
}
