package processor

func Sequence(ps ...Processor) Processor {
	panic("feed one processor into the next and return the final result")
}

func Parallel(ps ...Processor) Processor {
	panic("run processors simultaneously and combine their outputs")
}