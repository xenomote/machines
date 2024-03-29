package dfa

import (
	"github.com/xenomote/machine/assert"
	"github.com/xenomote/machine/event"
)

type dfa struct {
	current string
	states  map[string]map[string]transition
	inputs  []string
}

type transition struct {
	state  string
	output evt
}

func (m *dfa) State() string {
	return m.current
}

func (m *dfa) Step(received evt) (out evt) {
	defer func() { 
		if out == nil {
			assert.That("event cannot be nil")
		}
	}()

	e := event.Of()

	// if the machine is already out of bounds, do nothing
	s, exists := m.states[m.current]
	if !exists {
		return e
	}

	matched := false
	for i, t := range s {
		if !received.Matches(i) {
			continue
		}

		// matching multiple transitions is invalid behaviour for a deterministic machine
		if matched {
			assert.That("state", s, "event", received, "cannot match multiple transtions")
		}

		e = t.output

		matched = true
		m.current = t.state
	}

	// if the state had no transition for this event, go to the error state
	if !matched && m.couldMatch(received) {
		m.current = ""
	}

	return e
}

func (m *dfa) couldMatch(e evt) bool {
	for _, i := range m.inputs {
		if e.Matches(i) {
			return true
		}
	}
	return false
}
