# Lilt

Lilt is a little library (and tool) to copy input data to output on a regular interval

## Tool companion

Lilt command line tool copies stdin data to sdtout data on a regular interval.

### Installation

    $ go get github.com/squioc/lilt/cmd/lilt

### Usage

As a terminal,

    $ lilt [--period <period>]

with `<period>` the number of milliseconds between each line copy.

When Lilt is a terminal, it generates a new line at each period until SIGINT (CTRL+C) is catched.



As reading the output of another program,

    $ cat test.txt | lilt [--period <period>]

