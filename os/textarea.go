package main

import (
	"github.com/reiver/cyber80/os/log"
	"github.com/reiver/cyber80/os/webbed"

	"syscall/js"
)

var (
	textarea js.Value
)

func init() {
	textareas := webbed.Document.Call("getElementsByTagName", "textarea")

	if textareas.IsNull() {
		log.Panic("ERROR: ‘textareas’ is null")
		return
	}
	if textareas.IsUndefined() {
		log.Panic("ERROR: ‘textareas’ is undefined")
		return
	}

	if 1 > textareas.Length() {
		log.Panic("ERROR: no <textarea> found")
		return
	}

	textarea = textareas.Index(0)
}
