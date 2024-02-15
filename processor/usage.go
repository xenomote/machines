/*
Package processor defines processors and methods for composing them
*/
package processor

// Processor provides abstraction for sequential input processing
type Processor interface {
	State() string
	Step(Event) Event
}

// Event is a set of conditions which may be tested for
type Event interface {
	Matches(string) bool
	Empty() bool
}