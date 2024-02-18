# Event Processing Machines (`em`)

## Installation


## Usage



## Concepts

### Machine

machines are an abstraction of sequential event processing computations

machines may possess 

### Event

events are an opaque set of conditions 

events possess methods for testing if
- a given condition is set
- no conditions are set (the event is "empty")

### Condition

conditions are arbitrary strings which represent some true statement

conditions may overlap in meaning; for example given the following conditions
about a recieved character, and their respective meaning:
1. `[1-9]`: between 0 and 9 inclusive
2. `[01]`: equals 0 or 1
the first would be set for characters 2 to 9, the second for the character 0,
and both would be set for the character 1

### Deterministic Finite State Automaton (dfa)

todo

### Regular Expression (rgx)

todo

### Push Down Automaton (pda)

todo

### Context Free Grammars (cfg)

todo
