type ParkingSystem struct {
	places []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		places: []int{0, big, medium, small},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.places[carType] <= 0 {
		return false
	}

	this.places[carType]--
	return true
}
