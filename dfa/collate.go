package dfa

import (
	"fmt"

	"github.com/xenomote/machine/event"
)

func (c Config) collate() (*dfa, error) {
	if c.Start == "" {
		return nil, ErrNoStart
	}

	m := &dfa{
		current: c.Start,
		states:  map[string]map[string]transition{},
	}

	is := map[string]bool{}
	for n, s := range c.State {
		if _, exists := m.states[n]; exists {
			return nil, fail(ErrDuplicateState, n)
		}

		m.states[n] = map[string]transition{}

		for i, t := range s {
			if _, exists := m.states[n][i]; exists {
				return nil, fail(ErrDuplicateTransition, i, "in", n)
			}
			is[i] = true

			if len(t) < 1 {
				return nil, fail(ErrMissingTransition, "in", n)
			}

			m.states[n][i] = transition{
				state:  t[0],
				output: event.Of(t[1:]...),
			}
		}
	}

	for i := range is {
		m.inputs = append(m.inputs, i)
	}

	return m, nil
}

func fail(e error, msg ...any) (out error) {
	return fmt.Errorf("%w: %s", e, fmt.Sprint(msg...))
}
