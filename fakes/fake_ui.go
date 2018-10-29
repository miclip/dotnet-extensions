package fakes

import (
	"io"

	"github.com/miclip/dotnet-extensions/ui"
)

// NewFakeUI returns a fake UI
func NewFakeUI(outWriter, errWriter io.Writer, inReader io.Reader) *ui.WriterUI {
	return ui.NewWriterUI(outWriter, errWriter, inReader)
}
