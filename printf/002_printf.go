package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	// Stdout points to the output buffer to send screen output
	orgStdout io.Writer = os.Stdout
	// Stderr points to the output buffer to send errors to the screen
	orgStderr io.Writer = os.Stderr
)

func main() {
	oldStdout := orgStdout
	oldStderr := orgStderr

	// Create new buffers
	uStdout := new(bytes.Buffer)
	uStderr := new(bytes.Buffer)

	orgStdout = uStdout
	orgStderr = uStderr

	_ = uStderr
	//fmt.Printf("in Test Print Yaml")
	fmt.Fprintf(orgStdout, "in Test Print Yaml")
	lines := strings.Split(uStdout.String(), " ")

	fmt.Printf(" \n out %s  ", lines)
	orgStdout = oldStdout
	orgStderr = oldStderr
	fmt.Printf(" back to org ")
}
