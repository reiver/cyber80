package main

import (
	"fmt"
	"reflect"
)

func init() {
	if nil == exports["fmt"] {
		exports["fmt"] = map[string]reflect.Value{}
	}

	exports["fmt"]["Print"]   = reflect.ValueOf(export_fmt_Print)
	exports["fmt"]["Printf"]  = reflect.ValueOf(export_fmt_Printf)
	exports["fmt"]["Println"] = reflect.ValueOf(export_fmt_Println)
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
