package _1232

import "math"

func checkStraightLine(coordinates [][]int) bool {
	x1, y1 := coordinates[0][0], coordinates[0][1]
	x2, y2 := coordinates[1][0], coordinates[1][1]
	var pre float64
	if x2-x1 == 0 {
		pre = math.MaxFloat64
	} else {
		pre = (float64)(y2 - y1) / (float64)(x2 - x1)
	}

	for i := 2; i < len(coordinates); i++ {
		x1, y1 = coordinates[i-1][0], coordinates[i-1][1]
		x, y := coordinates[i][0], coordinates[i][1]
		var cur float64
		if x-x1 == 0 {
			cur = math.MaxFloat64
		} else {
			cur = (float64)(y - y1) / (float64)(x - x1)
		}
		if cur != pre {
			return false
		}
	}
	return true
}

