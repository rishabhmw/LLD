package parking_system

import "sync"

// TODO: can make this singleton

type IParkingLot interface {
	GetID() int
	Park(v IVehicle, at int)
	Vacate(at int)
	GetOccupancies() ([]bool, int)
	GetSpaces() ([]IParkingSpace, int)
	AcquireLock()
	ReleaseLock()
}

type SimpleParkingLot struct {
	id          int
	occupancies []bool
	spaces      []IParkingSpace
	floors      int
	mu          *sync.Mutex
}

func (s SimpleParkingLot) GetID() int {
	return s.id
}

func (s SimpleParkingLot) Park(v IVehicle, at int) {
	s.spaces[at].Occupy(v)
	s.occupancies[at] = true
}

func (s SimpleParkingLot) Vacate(at int) {
	s.spaces[at].Vacate()
	s.occupancies[at] = false
}

func (s SimpleParkingLot) GetOccupancies() ([]bool, int) {
	return s.occupancies, s.floors
}

func (s SimpleParkingLot) GetSpaces() ([]IParkingSpace, int) {
	return s.spaces, s.floors
}

func (s SimpleParkingLot) AcquireLock() {
	s.mu.Lock()
}

func (s SimpleParkingLot) ReleaseLock() {
	s.mu.Unlock()
}

func NewSimpleParkingLot(id int, floors int, numOfLots int) IParkingLot {
	size := floors * numOfLots
	spaces := make([]IParkingSpace, size)
	for i := 0; i < size; i++ {
		spaces[i] = NewSimpleParkingSpace(i, i%numOfLots, []string{
			VehicleTypeBike, VehicleTypeCar, VehicleTypeTruck,
		})
	}

	return &SimpleParkingLot{
		id:          id,
		spaces:      spaces,
		occupancies: make([]bool, size),
	}
}
