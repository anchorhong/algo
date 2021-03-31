package _1109

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_corpFlightBookings(t *testing.T) {
	tests := []struct {
		name       string
		bookingStr string
		n          int
		wantStr    string
	}{
		{
			name:       "testCase 1",
			bookingStr: "[[1,2,10],[2,3,20],[2,5,25]]",
			n:          5,
			wantStr:    "[10,55,45,25,25]",
		},
		{
			name:       "testCase 2",
			bookingStr: "[[1,2,10],[2,2,15]]",
			n:          2,
			wantStr:    "[10,25]",
		},
	}
	for _, tt := range tests {
		var bookings [][]int
		json.Unmarshal([]byte(tt.bookingStr), &bookings)
		var want []int
		json.Unmarshal([]byte(tt.wantStr), &want)

		t.Run(tt.name, func(t *testing.T) {
			if got := corpFlightBookings(bookings, tt.n); !reflect.DeepEqual(got, want) {
				t.Errorf("corpFlightBookings() = %v, want %v", got, want)
			}
		})
	}
}
