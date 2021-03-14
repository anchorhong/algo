package _989

func addToArrayForm(A []int, K int) []int {
	var stack []int
	carry := 0
	i := len(A) - 1
	for K != 0 || carry > 0 || i >= 0 {
		x := K % 10
		if i >= 0 {
			x += A[i]
			i--
		}
		x += carry
		carry = x / 10
		stack = append(stack, x%10)
		K /= 10
	}
	var res []int
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		res = append(res, top)
		stack = stack[:len(stack)-1]
	}
	return res
}
