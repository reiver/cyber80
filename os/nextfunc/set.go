package nextfunc

import (
	"github.com/reiver/cyber80/os/webbed"

	"syscall/js"
)

func Set(fn func()) error {
	wrapped := wrap(fn)

	jsFunc := js.FuncOf(wrapped)
	webbed.Window.Set(magicFunctionName, jsFunc)

	return nil
}
