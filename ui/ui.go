package ui

import (
	"fmt"
	"bufio"
	"io"
	"os"

)

//WriterUI Writer UI type
type WriterUI struct {
	outWriter    io.Writer
	errWriter    io.Writer
	inReader     *bufio.Reader
}

// NewConsoleUI creates an instance of the WriterUI
func NewConsoleUI() *WriterUI {
	return NewWriterUI(os.Stdout, os.Stderr, os.Stdin)
}

// NewWriterUI creates a UI instance
func NewWriterUI(outWriter, errWriter io.Writer, inReader io.Reader) *WriterUI {
	return &WriterUI{
		outWriter:    outWriter,
		errWriter:    errWriter,
		inReader:     bufio.NewReader(inReader),
	}
}

// PrintLinef requests input via stdin from the user
func (w *WriterUI) PrintLinef(pattern string, args ...interface{}) {
	fmt.Fprintf(w.outWriter, pattern+"\n", args...)
}

// Printf requests input via stdin from the user
func (w *WriterUI) Printf(pattern string, args ...interface{}) {
	fmt.Fprintf(w.outWriter, pattern, args...)
}

// ErrorLinef requests input via stdin from the user
func (w *WriterUI) ErrorLinef(pattern string, args ...interface{}) {
	fmt.Fprintf(w.errWriter, pattern+"\n", args...)
}

// Errorf requests input via stdin from the user
func (w *WriterUI) Errorf(pattern string, args ...interface{}) {
	fmt.Fprintf(w.errWriter, pattern, args...)
}


