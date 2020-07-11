package webbed

import (
	"syscall/js"
)

var (
	Window js.Value
)

func init() {
	Window = js.Global()

	if Window.IsNull() {
		msg := "ERROR: ‘window’ is null"

		panic(msg)
	}
	if Window.IsUndefined() {
		msg := "ERROR: ‘window’ is undefined"

		panic(msg)
	}
}
