/*
Package processor defines processors and methods for composing them

  - matches input events to state transitions, with optional output events
  - conditions not mentioned in transitions are a no-op to allow composability
  - states which are named but not defined treat every input as a no-op
  - inputs with no transition mappings go to the special empty error state
*/
package processor

// Processor provides abstraction for sequential event processing
type Processor interface {
	Step(Event) Event
}

// Processors which have some state may 
type Stateful interface {
	State() string
}

// Event is an opaque set of conditions which may be tested for
type Event interface {
	Matches(string) bool
	Empty() bool
}
