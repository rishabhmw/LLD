package parking_system

import "fmt"

type ITicket interface {
	Print()
	GetVehicle() IVehicle
	GetLot() IParkingLot
	GetParkedAt() (int, int)
}

type SimpleTicket struct {
	parkedAt int
	floor    int
	lot      IParkingLot

	vehicle IVehicle
}

func (t *SimpleTicket) Print() {
	fmt.Printf("Vechicle %d parked at lot %d on floor %d at space %d",
		t.vehicle.GetID(), t.lot.GetID(), t.floor, t.parkedAt)
}

func (t *SimpleTicket) GetVehicle() IVehicle {
	return t.vehicle
}

func (t *SimpleTicket) GetLot() IParkingLot {
	return t.lot
}

func (t *SimpleTicket) GetParkedAt() (int, int) {
	return t.parkedAt, t.floor
}
