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

	export.Symbols["fmt"]["Errorf"]   = reflect.ValueOf(fmt.Errorf)
	export.Symbols["fmt"]["Sprint"]   = reflect.ValueOf(fmt.Sprint)
	export.Symbols["fmt"]["Sprintf"]  = reflect.ValueOf(fmt.Sprintf)
	export.Symbols["fmt"]["Sprintln"] = reflect.ValueOf(fmt.Sprintln)
	export.Symbols["fmt"]["Print"]    = reflect.ValueOf(fmt.Print)
	export.Symbols["fmt"]["Printf"]   = reflect.ValueOf(fmt.Printf)
	export.Symbols["fmt"]["Println"]  = reflect.ValueOf(fmt.Println)
}
