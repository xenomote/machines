package event

import "strings"

type event map[string]bool

func (e event) Matches(i string) bool {
	return e[i]
}

func (e event) Empty() bool {
	return len(e) == 0
}

func (e event) String() string {
	s := strings.Builder{}

	s.WriteRune('(')

	sep := false
	for k := range e {
		if sep {
			s.WriteString(", ")
		} else {
			sep = true
		}
		s.WriteRune('"')
		s.WriteString(k)
		s.WriteRune('"')
	}
	s.WriteRune(')')

	return s.String()
}
