package main

import (
	"io"
	"os"
	"time"

	flag "github.com/spf13/pflag"
        "github.com/squioc/lilt/pkg/lilt"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	var period int
	flag.IntVar(&period, "period", 1000, "number of milliseconds between each pulse")
        flag.Parse()

	var reader io.Reader
	if terminal.IsTerminal(int(os.Stdin.Fd())) {
		reader = &lilt.LineFeedReader{}
	} else {
		reader = os.Stdin
	}

	tempo := lilt.NewLiltWithReader(reader, os.Stdout, time.NewTicker(time.Duration(period)*time.Millisecond))
	tempo.Run()

}
