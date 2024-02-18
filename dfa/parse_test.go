package dfa_test

func fromSource() pst {
	return fromConfig()

	// // once implemented
	// return dfa.Must(`
    // start
    //     [1-9]   int
    //     '-'     neg
    //     '0'     zer
        
    // int
    //     eof     success
    //     [0-9]   dig
    //     [eE]    sci
    //     '.'     dot
    
    // neg
    //     eof     success
    //     [1-9]   int
    //     '0'     zer
    
    // dig
    //     eof     success
    //     [0-9]   dig
    //     [eE]    sci
    //     '.'     dot
    
    // zer
    //     eof     success
    //     [eE]    sci
    //     '.'     dot
    
    // dot
    //     [0-9]   dec
    
    // dec
    //     eof     success
    //     [0-9]   dec
    //     [eE]    sci
    
    // sci
    //     [0-9]   exp
    //     [+-]    sgn
    
    // sgn
    //     [0-9]   exp
    
    // exp
    //     eof     success
    //     [0-9]   exp
    // `)
}

func ExampleMustParse() {
	driver(fromSource(), examples)
	// Output:
}