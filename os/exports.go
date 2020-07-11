package main

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
	// • exports["bufio"]
	// • exports["bytes"]
	// • exports["fmt"]
	// • exports["image"]
	// • exports["image/color"]
	// • exports["image/color/palette"]
	// • exports["math"]
	// • exports["math/big"]
	// • exports["math/rand"]
	// • exports["time"]
	// • etc etc etc
	//
	// But you could also include other (custom) packages too; for example?
	//
	// • exports["apple/banana/cherry"]
	// • exports["github.com/reiver/go-hashuri"]
	// • etc etc etc
	//
	// Underneath this top-level key is what is being exported from the package. For example:
	//
	// • exports["fmt"]["Errorf"]
	// • exports["fmt"]["Fprint"]
	// • exports["fmt"]["Fprintf"]
	// • exports["fmt"]["Fprintln"]
	// • exports["fmt"]["Fscan"]
	// • exports["fmt"]["Fscanf"]
	// • exports["fmt"]["Fscanln"]
	// • exports["fmt"]["Print"]
	// • exports["fmt"]["Printf"]
	// • exports["fmt"]["Println"]
	// • exports["fmt"]["Scan"]
	// • exports["fmt"]["Scanf"]
	// • exports["fmt"]["Scanln"]
	// • exports["fmt"]["Sprint"]
	// • exports["fmt"]["Sprintf"]
	// • exports["fmt"]["Sprintln"]
	// • exports["fmt"]["Sccan"]
	// • exports["fmt"]["Scanf"]
	// • exports["fmt"]["Scanln"]
	// • exports["fmt"]["Formatter"]
	// • exports["fmt"]["GoStringer"]
	// • exports["fmt"]["ScanState"]
	// • exports["fmt"]["Scanner"]
	// • exports["fmt"]["State"]
	// • exports["fmt"]["Stringer"]
	//
	// You could give each of these some implementation. For example:
	//
	//	func my_Print(a ...interface{}) (n int, err error) {
	//		// ...
	//	}
	//	
	//	exports["fmt"]["Print"] = reflect.ValueOf(my_Print)
	//
	// In this code, other files set the values of this package-scope variable.
	exports map[string]map[string]reflect.Value = map[string]map[string]reflect.Value{}
)
