package main

import (
	"github.com/reiver/cyber80/os/webbed"

	"syscall/js"
)

var (
	textarea js.Value
)

func init() {
	textareas := webbed.Document.Call("getElementsByTagName", "textarea")

	if textareas.IsNull() {
		msg := "ERROR: ‘textareas’ is null"

		log(msg)
		panic(msg)
	}
	if textareas.IsUndefined() {
		msg := "ERROR: ‘textareas’ is undefined"

		log(msg)
		panic(msg)
	}

	if 1 > textareas.Length() {
		msg := "ERROR: no <textarea> found"

		log(msg)
		panic(msg)
	}

	textarea = textareas.Index(0)
}
