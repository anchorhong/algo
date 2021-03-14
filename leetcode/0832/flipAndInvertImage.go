package _832

func flipAndInvertImage(A [][]int) [][]int {
	for _, row := range A {
		l, r := 0, len(row)-1
		for ; l < r; l, r = l+1, r-1 {
			if row[l] == row[r] {
				row[l] ^= 1
				row[r] ^= 1
			}
		}
		if l == r {
			row[l] ^= 1
		}
	}
	return A
}
