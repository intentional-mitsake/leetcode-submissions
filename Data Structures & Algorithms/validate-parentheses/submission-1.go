func isValid(s string) bool {
    // first bracket opened is always the last bracket closed
	// "()" or "[{()}]" or "{}()"--> LIFO structure, naturally stack
	var stack []rune
	bracketMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, c := range s {
		// if opening brakcet((, {, [), push
		if c == rune('(') || c == rune('[') || c == rune('{') {
			stack = append(stack, c)
		} else {
			// if the stack if empty, no point in cheking
			if len(stack) > 0{
				// check if this closing bracket is opened at ToS
				// [ { (--ToS-- --> true so pushed, now next bracket--> ) --> check if at ToS
				if bracketMap[c] != stack[len(stack)-1] {
					// if not closed at ToS, not valid, false
				    return false
			    }
			    // pop for next 
			    stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	// theres a case with only one char('[') in the string
	// for that, it will be pushed to stack but not popped
	// in other cases stack will be emptied if valid
	// also the case for ']' only in input is taken care of by if(len(stack)>0)
	return len(stack) == 0
}
