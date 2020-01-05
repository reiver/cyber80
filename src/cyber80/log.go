package main

import (
	"fmt"
	"strings"
)

func log(a ...interface{}) {

	var builder strings.Builder
	builder.WriteString(name)
	builder.WriteString(": ")
	builder.WriteString(fmt.Sprint(a...))
	builder.WriteString("\n")

	fmt.Print(builder.String())
}

func logf(format string, a ...interface{}) {

	var builder strings.Builder

	builder.WriteString(name)
	builder.WriteString(": ")
	builder.WriteString(format)
	builder.WriteString("\n")

	fmt.Printf(builder.String(), a...)
}
