package dfa_test

import (
	"fmt"

	"github.com/xenomote/machine/event"
	"github.com/xenomote/machine/processor"
)

type (
	evt = processor.Event
	pro = processor.Processor
)

// a battery of test examples
var examples = []string{
    "0",            // simplest passing case
    "123",          // multiple digits
    "100",          // multiple zeroes
    "-104",         // negative number
    "123.0341",     // decimal
    "1e4",          // exponent
    "10E33",        // multiple digit exponent
    "0e-4",         // negative exponent
    "534.23E+23",   // decimal, exponent, optional sign
}

// mapper from characters to relevant (non-mutually exclusive) conditions
func mapper(c rune) evt {
	return event.Where(map[string]bool{
		"[1-9]": '1' <= c && c <= '9',
		"[0-9]": '0' <= c && c <= '9',
		"[eE]":  c == 'e' || c == 'E',
		"[+-]":  c == '+' || c == '-',
		"0":     c == '0',
		"-":     c == '-',
		".":     c == '.',
	})
}

// runs the provided machine over each example until an error is found
func driver(m pro, examples []string) {
    for n, example := range examples {
        for i, c := range example {
            m.Step(mapper(c))
            if m.State() == "" {
                pln("example", n+1, "error at position ", i)
            }
        }
    
        m.Step(event.Of("eof"))
        if m.State() != "success" {
            pln("example", n+1, "failed to match")
        }
	}
}

// helper for panic with a nice message
func pln(msg ...any) {
    panic(fmt.Sprintln(msg...))
}