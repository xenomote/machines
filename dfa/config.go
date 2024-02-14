package dfa

// data type for configuring DFAs directly
type Config struct {
	Start string
	State ConfigState
}

type ConfigState = map[string]map[string][]string