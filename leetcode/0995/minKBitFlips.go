package _995

func minKBitFlips(A []int, K int) (ans int) {
	n := len(A)
	diff := make([]int, n+1)
	revCnt := 0
	for i, v := range A {
		revCnt ^= diff[i]
		if v^revCnt == 0 {
			if i+K > n {
				return -1
			}
			ans++
			revCnt ^= 1
			diff[i+K] ^= 1
		}
	}
	return
}

func minKBitFlipsOps(A []int, K int) (ans int) {
	n := len(A)
	revCnt := 0
	for i, v := range A {
		if i >= K && A[i-K] > 1 {
			revCnt ^= 1
			A[i-K] -= 2 // 复原数组元素，若允许修改数组 A，则可以省略
		}
		if v == revCnt {
			if i+K > n {
				return -1
			}
			ans++
			revCnt ^= 1
			A[i] += 2
		}
	}
	return
}

func minKBitFlipsQueue(A []int, K int) int {
	var queue []int
	var ans int
	for i := range A {
		if len(queue) > 0 && queue[0]+K <= i {
			queue = queue[1:]
		}
		if len(queue)%2 == A[i] {
			queue = append(queue, i)
			ans++
		}
	}
	return ans
}
