package dfa

import (
	"errors"

	"github.com/xenomote/machine/bug"
	"github.com/xenomote/machine/event"
	"github.com/xenomote/machine/processor"
)

// package internal alias for brevity
type (
	evt = event.Event
	pro = processor.Processor
)

func Must(src string) (out pro) {
	return MustParse(src).MustCollate()
}

func MustParse(s string) *Config {
	c, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return c
}

func (c Config) MustCollate() *dfa {
	m, err := c.Collate()
	if err != nil {
		panic(err)
	}
	return m
}

func Parse(src string) (out *Config, err error) {
	out, err = parse(src)

	if (out == nil) && (err == nil) {
		bug.Exit("config and error cannot not both be nil")
	}

	if (out != nil) && (err != nil) {
		bug.Exit("config and error cannot not both be non-nil:", err)
	}
	
	return
}

var (
	ErrNoStart             = errors.New("no start state")
	ErrDuplicateState      = errors.New("duplicated state")
	ErrDuplicateTransition = errors.New("duplicated transition")
	ErrTransitionToEmpty   = errors.New("invalid transition to empty state")
	ErrMissingTransition   = errors.New("missing transition input")
)

func (c Config) Collate() (dfa *dfa, err error) {
	dfa, err = c.collate()

	if (dfa == nil) && (err == nil) {
		bug.Exit("dfa and error cannot both be nil")
	}

	if (dfa != nil) && (err != nil) {
		bug.Exit("dfa and error cannot both be non-nil:", err)
	}

	if dfa == nil {
		return dfa, err
	}

	if _, exists := dfa.states[""]; exists {
		bug.Exit("empty state cannot be present in the states map")
	}

	for n, s := range dfa.states {
		for i, t := range s {
			if t.state == "" {
				bug.Exit("empty state cannot be present in transitions state", n, i)
			}

			if t.output == nil {
				bug.Exit("nil event cannot be present in the transitions output", n, i)
			}
		}
	}

	return
}
