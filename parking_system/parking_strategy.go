package parking_system

import "fmt"

type IParkingStrategy interface {
	Execute(lot IParkingLot, v IVehicle) (int, int, error)
}

type SimpleParkingStrategy struct {
}

func (s SimpleParkingStrategy) Execute(lot IParkingLot, v IVehicle) (int, int, error) {
	lot.AcquireLock()
	defer lot.ReleaseLock()

	spaces, floor := lot.GetSpaces()
	occupancies, _ := lot.GetOccupancies()

	var parkAt int
	var found bool
	for i, occupancy := range occupancies {
		if !occupancy && spaces[i].DoesSupport(v.GetType()) {
			parkAt = i
			found = true
			break
		}
	}
	if !found {
		return 0, 0, fmt.Errorf("no suitable parking spot found")
	}

	lot.Park(v, parkAt)
	return parkAt, floor, nil
}

func NewSimpleParkingStrategy() IParkingStrategy {
	return &SimpleParkingStrategy{}
}
