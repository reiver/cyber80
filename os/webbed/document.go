package webbed

import (
	"syscall/js"
)

var (
	Document js.Value
)

func init() {
	Document = js.Global().Get("document")

	if Document.IsNull() {
		msg := "ERROR: ‘document’ is null"

		panic(msg)
	}
	if Document.IsUndefined() {
		msg := "ERROR: ‘document’ is undefined"

		panic(msg)
	}
}
