package verboten

import (
	"github.com/reiver/cyber80/os/export"

	"fmt"
	"reflect"
)

func init() {
	if nil == export.Symbols["fmt"] {
		export.Symbols["fmt"] = map[string]reflect.Value{}
	}

	export.Symbols["fmt"]["Print"]   = reflect.ValueOf(export_fmt_Print)
	export.Symbols["fmt"]["Printf"]  = reflect.ValueOf(export_fmt_Printf)
	export.Symbols["fmt"]["Println"] = reflect.ValueOf(export_fmt_Println)
}

func export_fmt_Print(a ...interface{}) (n int, err error) {
	return fmt.Print(a...)
}

func export_fmt_Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}

func export_fmt_Println(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}
