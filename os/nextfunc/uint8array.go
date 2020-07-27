package nextfunc

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
		msg := "ERROR: ‘uint8Array’ is null"

		log.Publish(msg)
		panic(msg)
	}
	if uint8Array.IsUndefined() {
		msg := "ERROR: ‘uint8Array’ is undefined"

		log.Publish(msg)
		panic(msg)
	}
}
