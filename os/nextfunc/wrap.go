package nextfunc

import (
	"github.com/reiver/go-c80"

	"syscall/js"
)

func wrap(fn func()) func(this js.Value, args []js.Value) interface{} {

	next := func(this js.Value, args []js.Value) interface{} {

		fn()

		u8array := uint8Array.New(args[0])
		_ = js.CopyBytesToJS(u8array, c80.Frame())

		return nil
	}

	return next
}
