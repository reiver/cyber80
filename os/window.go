package main

import (
	"syscall/js"
)

var (
	window js.Value
)

func init() {
	window = js.Global()

	if window.IsNull() {
		msg := "ERROR: ‘window’ is null"

		log(msg)
		panic(msg)
	}
	if window.IsUndefined() {
		msg := "ERROR: ‘window’ is undefined"

		log(msg)
		panic(msg)
	}
}
