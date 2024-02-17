/*
Package dfa defines deterministic finite state automata processors
*/
package dfa

import (
	"errors"

	"github.com/xenomote/machine/assert"
	"github.com/xenomote/machine/processor"
)

var (
	ErrNoStart             = errors.New("no start state")
	ErrDuplicateState      = errors.New("duplicated state")
	ErrDuplicateTransition = errors.New("duplicated transition")
	ErrTransitionToEmpty   = errors.New("invalid transition to empty state")
	ErrMissingTransition   = errors.New("missing transition input")
)

// package internal alias for brevity
type (
	evt = processor.Event
	pro = processor.Processor
)

// Must is a helper method for generating a DFA directly from known-good code
// this method will panic if any errors are found
func Must(src string) (out pro) {
	return MustParse(src).MustCollate()
}

// MustParse is a helper method for parsing DFA configs from known-good code
// this method will panic if any errors are found
func MustParse(s string) *Config {
	c, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return c
}

// MustCollate is a helper method for creating a new DFA instance from config
// this method will panic if any errors are found
func (c Config) MustCollate() *dfa {
	m, err := c.Collate()
	if err != nil {
		panic(err)
	}
	return m
}

// Parse parses DFA config from source
// an error is returned if the source is invalid
func Parse(src string) (out *Config, err error) {
	out, err = parse(src)

	if (out == nil) && (err == nil) {
		assert.That("config and error cannot not both be nil")
	}

	if (out != nil) && (err != nil) {
		assert.That("config and error cannot not both be non-nil:", err)
	}

	return
}

// Collate creates a new DFA instance from config
// an error is returned if the config is invalid
func (c Config) Collate() (dfa *dfa, err error) {
	dfa, err = c.collate()

	if (dfa == nil) && (err == nil) {
		assert.That("dfa and error cannot both be nil")
	}

	if (dfa != nil) && (err != nil) {
		assert.That("dfa and error cannot both be non-nil:", err)
	}

	if dfa == nil {
		return dfa, err
	}

	if _, exists := dfa.states[""]; exists {
		assert.That("empty state cannot be present in the states map")
	}

	for n, s := range dfa.states {
		for i, t := range s {
			if t.state == "" {
				assert.That("empty state cannot be present in transitions state", n, i)
			}

			if t.output == nil {
				assert.That("nil event cannot be present in the transitions output", n, i)
			}
		}
	}

	return
}
