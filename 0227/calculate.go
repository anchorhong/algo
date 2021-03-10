package _227

func calculate(s string) (ans int) {
	var stack []int
	defaultSign := '+'
	num := 0
	for i, ch := range s {
		isDigit := ch >= '0' && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch defaultSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			defaultSign = ch
			num = 0
		}
	}
	for _, s := range stack {
		ans += s
	}
	return
}
