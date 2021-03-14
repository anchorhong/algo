package _338

func countBits(num int) []int {
	bits := make([]int, num+1)
	highBit := 0
	for i := 1; i <= num; i++ {
		if i & (i-1) == 0 {
			highBit = i
		}
		bits[i] = bits[i-highBit] + 1
	}
	return bits
}
