package _888

func fairCandySwap(A []int, B []int) []int {
	var aSum int
	var bSum int
	bValue := make(map[int]interface{})
	for _, a := range A {
		aSum += a
	}
	for _, b := range B {
		bSum += b
		bValue[b] = nil
	}
	diff := (aSum - bSum) / 2
	for _, a := range A {
		y := a - diff
		if _, ok := bValue[y]; ok {
			return []int{a, y}
		}
	}
	return nil
}
