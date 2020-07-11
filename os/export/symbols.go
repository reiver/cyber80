package export

import (
	"reflect"
)

var (
	// This is used for the Go interpreter.
	//
	// This lets us add in functionality, in the form of packages, and the things those
	// packages export.
	//
	// The top level keys are the package names.
	//
	// For example, they could be built-in package names such as:
	//
	// • export.Symbols["bufio"]
	// • export.Symbols["bytes"]
	// • export.Symbols["fmt"]
	// • export.Symbols["image"]
	// • export.Symbols["image/color"]
	// • export.Symbols["image/color/palette"]
	// • export.Symbols["math"]
	// • export.Symbols["math/big"]
	// • export.Symbols["math/rand"]
	// • export.Symbols["time"]
	// • etc etc etc
	//
	// But you could also include other (custom) packages too; for example?
	//
	// • export.Symbols["apple/banana/cherry"]
	// • export.Symbols["github.com/reiver/go-hashuri"]
	// • etc etc etc
	//
	// Underneath this top-level key is what is being exported from the package. For example:
	//
	// • export.Symbols["fmt"]["Errorf"]
	// • export.Symbols["fmt"]["Fprint"]
	// • export.Symbols["fmt"]["Fprintf"]
	// • export.Symbols["fmt"]["Fprintln"]
	// • export.Symbols["fmt"]["Fscan"]
	// • export.Symbols["fmt"]["Fscanf"]
	// • export.Symbols["fmt"]["Fscanln"]
	// • export.Symbols["fmt"]["Print"]
	// • export.Symbols["fmt"]["Printf"]
	// • export.Symbols["fmt"]["Println"]
	// • export.Symbols["fmt"]["Scan"]
	// • export.Symbols["fmt"]["Scanf"]
	// • export.Symbols["fmt"]["Scanln"]
	// • export.Symbols["fmt"]["Sprint"]
	// • export.Symbols["fmt"]["Sprintf"]
	// • export.Symbols["fmt"]["Sprintln"]
	// • export.Symbols["fmt"]["Sccan"]
	// • export.Symbols["fmt"]["Scanf"]
	// • export.Symbols["fmt"]["Scanln"]
	// • export.Symbols["fmt"]["Formatter"]
	// • export.Symbols["fmt"]["GoStringer"]
	// • export.Symbols["fmt"]["ScanState"]
	// • export.Symbols["fmt"]["Scanner"]
	// • export.Symbols["fmt"]["State"]
	// • export.Symbols["fmt"]["Stringer"]
	//
	// You could give each of these some implementation. For example:
	//
	//	func my_Print(a ...interface{}) (n int, err error) {
	//		// ...
	//	}
	//	
	//	export.Symbols["fmt"]["Print"] = reflect.ValueOf(my_Print)
	//
	// In this code, other files set the values of this package-scope variable.
	Symbols map[string]map[string]reflect.Value = map[string]map[string]reflect.Value{}
)
