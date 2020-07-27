package main

import (
	"github.com/reiver/cyber80/os/log"

	"syscall/js"
)

var (
	uint8Array js.Value
)

func init() {
	uint8Array = js.Global().Get("Uint8Array")

	if uint8Array.IsNull() {
		log.Panic("ERROR: ‘uint8Array’ is null")
		return
	}
	if uint8Array.IsUndefined() {
		log.Panic("ERROR: ‘uint8Array’ is undefined")
		return
	}
}
