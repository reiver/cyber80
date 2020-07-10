package main

import (
	"syscall/js"
)

var (
	uint8Array js.Value
)

func init() {
	uint8Array = js.Global().Get("Uint8Array")

	if uint8Array.IsNull() {
		msg := "ERROR: ‘uint8Array’ is null"

		log(msg)
		panic(msg)
	}
	if uint8Array.IsUndefined() {
		msg := "ERROR: ‘uint8Array’ is undefined"

		log(msg)
		panic(msg)
	}
}
