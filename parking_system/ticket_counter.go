package parking_system

import "fmt"

type ITicketCounter interface {
	GetID() int
	GetParkingLot() IParkingLot
	GetTicket(v IVehicle, using IParkingStrategy) (ITicket, error)
	MarkExit(t ITicket)
}

type SimpleTicketCounter struct {
	id  int
	lot IParkingLot
}

func (s SimpleTicketCounter) GetID() int {
	return s.id
}

func (s SimpleTicketCounter) GetParkingLot() IParkingLot {
	return s.lot
}

func (s SimpleTicketCounter) GetTicket(v IVehicle, using IParkingStrategy) (ITicket, error) {
	parkedAt, floor, err := using.Execute(s.lot, v)
	if err != nil {
		return nil, fmt.Errorf("error getting ticket: %w", err)
	}
	return &SimpleTicket{
		parkedAt: parkedAt,
		floor:    floor,
		lot:      s.lot,
		vehicle:  v,
	}, err
}

func (s SimpleTicketCounter) MarkExit(t ITicket) {
	lot := t.GetLot()
	lot.AcquireLock()
	defer lot.ReleaseLock()

	parkedAt, _ := t.GetParkedAt()
	lot.Vacate(parkedAt)
}

func NewSimpleTicketCounter(id int, lot IParkingLot) ITicketCounter {
	return &SimpleTicketCounter{
		id:  id,
		lot: lot,
	}
}
