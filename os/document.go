package main

import (
	"syscall/js"
)

var (
	document js.Value
)

func init() {
	document = js.Global().Get("document")

	if document.IsNull() {
		msg := "ERROR: ‘document’ is null"

		log(msg)
		panic(msg)
	}
	if document.IsUndefined() {
		msg := "ERROR: ‘document’ is undefined"

		log(msg)
		panic(msg)
	}
}
