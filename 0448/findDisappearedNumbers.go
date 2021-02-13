package _448

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for _, num := range nums {
		v := (num - 1) % n
		nums[v] += n
	}

	var res []int
	for i, num := range nums {
		if num <= n {
			res = append(res, i+1)
		}
	}
	return res
}
