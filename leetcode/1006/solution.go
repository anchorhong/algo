package _1006

func clumsy(N int) int {
	if N <= 2 {
		return N
	}
	var res int
	res = N * (N - 1) / (N - 2)
	N -= 3
	times := N / 4
	surplus := N % 4
	for i := 0; i < times; i++ {
		res += N - (N-1)*(N-2)/(N-3)
		N -= 4
	}
	if surplus != 0 {
		res += 1
	}
	return res
}
