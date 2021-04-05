package _088

func merge(nums1 []int, m int, nums2 []int, n int) {
	var res []int
	for i, j := 0, 0; i <= m && j <= n; {
		if i == m {
			res = append(res, nums2[j:]...)
			i++
		} else if j == n {
			res = append(res, nums1[i:]...)
			j++
		} else if nums1[i] <= nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	copy(nums1, res)
}
