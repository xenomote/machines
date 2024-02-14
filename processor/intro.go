package processor

import "github.com/xenomote/machine/event"

// package internal alias for brevity
type evt = event.Event

type Processor interface {
	State() string
	Step(evt) evt
}

func Pipeline(ps ...Processor) Processor {
	panic("feed one processor into the next")
}
