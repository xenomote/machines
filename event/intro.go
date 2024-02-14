package event

import "github.com/xenomote/machine/bug"

type Event interface {
	Matches(string) bool
	Empty() bool
}

func check(e Event) {
	if e == nil {
		bug.Exit("event cannot be nil")
	}

	if e.Matches("") {
		bug.Exit("events cannot match empty condition")
	}
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
