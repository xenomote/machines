package dfa_test

import (
	"fmt"

	"github.com/xenomote/machine/dfa"
	"github.com/xenomote/machine/event"
	"github.com/xenomote/machine/processor"
)

type (
    evt = processor.Event
    pro = processor.Processor
)

func reset() pro {
    return dfa.Must(`
    start
        [1-9]   int
        '-'     neg
        '0'     zer
        
    int
        eof     success
        [0-9]   dig
        [eE]    sci
        '.'     dot
    
    neg
        eof     success
        [1-9]   int
        '0'     zer
    
    dig
        eof     success
        [0-9]   dig
        [eE]    sci
        '.'     dot
    
    zer
        eof     success
        [eE]    sci
        '.'     dot
    
    dot
        [0-9]   dec
    
    dec
        eof     success
        [0-9]   dec
        [eE]    sci
    
    sci
        [0-9]   exp
        [+-]    sgn
    
    sgn
        [0-9]   exp
    
    exp
        eof     success
        [0-9]   exp
    `)
}

func driver(c rune) evt {
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

func Example_succeeds() {
next:
	for n, example := range []string{
		"0",
		"123",
		"100",
		"1e4",
		"10E33",
        "0e-4",
        "534.23E+23",
		"-104",
		"1234.03401",
	} {
        m := reset()

		for i, c := range example {
			m.Step(driver(c))
			if m.State() == "" {
				fmt.Println(n+1, "error at position ", i)
				continue next
			}
		}

		m.Step(event.Of("eof"))
		if m.State() != "success" {
			fmt.Println(n+1, "failed to match")
			continue next
		}
	}

    // Output:
}