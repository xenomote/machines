/*
Package event defines utility methods for creating events
*/
package event

import (
	"github.com/xenomote/machine/assert"
	"github.com/xenomote/machine/processor"
)

type evt = processor.Event

func Of(cs ...string) evt {
	e := event{}
	for _, c := range cs {
		e[c] = true
	}

	check(e)
	return e
}

func Where(conditions map[string]bool) evt {
	e := event{}
	for c, exists := range conditions {
		if exists {
			e[c] = true
		}
	}

	check(e)
	return e
}

func check(e evt) {
	if e == nil {
		assert.That("event cannot be nil")
	}

	if e.Matches("") {
		assert.That("events cannot match empty condition")
	}
}
