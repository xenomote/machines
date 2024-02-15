/*
Package event defines the events consumed and produced by processors
*/
package event

import "github.com/xenomote/machine/assert"

type Event interface {
	Matches(string) bool
	Empty() bool
}

func Of(cs ...string) Event {
	e := evt{}
	for _, c := range cs {
		e[c] = true
	}

	check(e)
	return e
}

func Where(conditions map[string]bool) Event {
	e := evt{}
	for c, exists := range conditions {
		if exists {
			e[c] = true
		}
	}

	check(e)
	return e
}

func check(e Event) {
	if e == nil {
		assert.That("event cannot be nil")
	}

	if e.Matches("") {
		assert.That("events cannot match empty condition")
	}
}
