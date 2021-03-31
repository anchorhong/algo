package _1109

func corpFlightBookings(bookings [][]int, n int) []int {
	count := make([]int, n)
	for _, booking := range bookings {
		count[booking[0]-1] += booking[2]
		if booking[1] < n {
			count[booking[1]] -= booking[2]
		}
	}

	for i := 1; i < n; i++ {
		count[i] += count[i-1]
	}
	return count
}
