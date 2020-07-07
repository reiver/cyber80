package main

import (
	"fmt"
	"strings"
)

// log outputs a log message.
//
// If magic256 is run inside a web browser then the log will appear in the web browser's developer console.
// I.e., this is similar to JavaScript's console.log().
//
// Example
//
// If the following is executed:
//
//	log("Hello world!")
//
// Then the output will be:
//
//	magic256: Hello world!
func log(a ...interface{}) {

	var builder strings.Builder
	builder.WriteString(name)
	builder.WriteString(": ")
	builder.WriteString(fmt.Sprint(a...))
	builder.WriteString("\n")

	fmt.Print(builder.String())
}

// logf outputs a formatted log message.
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
//	logf("Hello %s!", name)
//
// Then the output will be:
//
//	magic256: Hello Joe Blow!
func logf(format string, a ...interface{}) {

	var builder strings.Builder

	builder.WriteString(name)
	builder.WriteString(": ")
	builder.WriteString(format)
	builder.WriteString("\n")

	fmt.Printf(builder.String(), a...)
}
