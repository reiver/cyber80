package log

import (
	"fmt"
	"strings"
)

// Publish outputs a log message.
//
// If magic256 is run inside a web browser then the log will appear in the web browser's developer console.
// I.e., this is similar to JavaScript's console.log().
//
// Example
//
// If the following is executed:
//
//	Publish("Hello world!")
//
// Then the output will be:
//
//	magic256: Hello world!
func Publish(a ...interface{}) {

	var builder strings.Builder
	builder.WriteString(name)
	builder.WriteString(": ")
	builder.WriteString(fmt.Sprint(a...))
	builder.WriteString("\n")

	fmt.Print(builder.String())
}

// Publishf outputs a formatted log message.
//
// If magic256 is run inside a web browser then the log will appear in the web browser's developer console.
// I.e., this is similar to JavaScript's console.log().
//
// Example
//
// If the following is executed:
//
//	var name = "Joe Blow"
//
//	Publishf("Hello %s!", name)
//
// Then the output will be:
//
//	magic256: Hello Joe Blow!
func Publishf(format string, a ...interface{}) {

	var builder strings.Builder

	builder.WriteString(name)
	builder.WriteString(": ")
	builder.WriteString(format)
	builder.WriteString("\n")

	fmt.Printf(builder.String(), a...)
}
