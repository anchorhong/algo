package _1603

type ParkingSystem struct {
	park []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	park := make([]int, 3)
	park[0] = big
	park[1] = medium
	park[2] = small
	return ParkingSystem{park}
}

func (p *ParkingSystem) AddCar(carType int) bool {
	if p.park[carType-1] > 0 {
		p.park[carType-1]--
		return true
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
