func evalRPN(tokens []string) int {
	stack := make([]int, 0) // to store nums
	for _, el := range tokens{
		switch el {
			case "+":
			// pop two nums
			  a := stack[len(stack)-1]
			  stack = stack[:len(stack)-1]
			  b := stack[len(stack)-1]
			  stack = stack[:len(stack)-1]
			  res := a + b
			  // store res at tos for next operand pop
			  stack = append(stack, res)
			case "-":
			  a := stack[len(stack)-1]
			  stack = stack[:len(stack)-1]
			  b := stack[len(stack)-1]
			  stack = stack[:len(stack)-1]
			  res := b - a // tos el is subtracted
			  stack = append(stack, res)
			case "*":
		      a := stack[len(stack)-1]
			  stack = stack[:len(stack)-1]
			  b := stack[len(stack)-1]
			  stack = stack[:len(stack)-1]
			  res := a * b
			  stack = append(stack, res)
			case "/":
			  a := stack[len(stack)-1]
			  stack = stack[:len(stack)-1]
			  b := stack[len(stack)-1]
			  stack = stack[:len(stack)-1]
			  res := b / a // tos is divisor
			  stack = append(stack, res)
			default:
			  num, _ := strconv.Atoi(el)
			  stack = append(stack, num)
		}
	}
	// result is at tos after for loop
	return stack[len(stack)-1]

}
