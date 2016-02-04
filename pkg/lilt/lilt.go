package lilt

import (
	"bufio"
	"io"
	"time"
)

// Lilt provides copy bytes from the input to the output according to a period
type Lilt struct {
	input  *bufio.Scanner
	output io.Writer
	pulse  *time.Ticker
}

// NewLilt creates a new lilt
func NewLilt(input *bufio.Scanner, output io.Writer, pulse *time.Ticker) *Lilt {
	return &Lilt{input: input, output: output, pulse: pulse}
}

// NewLiltWithReader creates a new lilt with a reader as input
func NewLiltWithReader(reader io.Reader, output io.Writer, pulse *time.Ticker) *Lilt {
	return &Lilt{input: bufio.NewScanner(reader), output: output, pulse: pulse}
}

// Run copies bytes from the input to the output
func (l *Lilt) Run() error {
	for {
		if !l.input.Scan() {
			return nil
		}
		l.output.Write(l.input.Bytes())
		l.output.Write([]byte{0x0A})
		<-l.pulse.C
	}
	return nil
}
