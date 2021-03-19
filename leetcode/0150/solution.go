package _150

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int
	for _, s := range tokens {
		if s == "+" || s == "-" || s == "*" || s == "/" {
			op1, op2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[0 : len(stack)-2]
			switch s {
			case "+":
				stack = append(stack, op1+op2)
			case "-":
				stack = append(stack, op1-op2)
			case "*":
				stack = append(stack, op1*op2)
			case "/":
				stack = append(stack, op1/op2)
			}
		} else {
			r, _ := strconv.Atoi(s)
			stack = append(stack, r)
		}
	}
	return stack[0]
}
