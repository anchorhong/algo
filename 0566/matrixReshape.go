package _566

func matrixReshape(nums [][]int, r int, c int) [][]int {
	originM, originN := len(nums), len(nums[0])
	if originN*originM != r*c || (originM == r && originN == c) {
		return nums
	}
	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
	}
	for i := 0; i < originM*originN; i++ {
		res[i/r][i%c] = nums[i/originN][i%originN]
	}
	return res
}
