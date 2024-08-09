package parking_system

type IParkingSpace interface {
	GetID() int
	GetFloor() int
	Occupy(v IVehicle)
	Vacate()
	IsOccupied() (bool, IVehicle)
	SupportedVehiclesTypes() map[string]bool
	DoesSupport(vehicleType string) bool
}

type SimpleParkingSpace struct {
	id                     int
	floor                  int
	vehicle                *IVehicle
	supportedVehiclesTypes map[string]bool
}

func (s SimpleParkingSpace) GetID() int {
	return s.id
}

func (s SimpleParkingSpace) GetFloor() int {
	return s.floor
}

func (s SimpleParkingSpace) Occupy(v IVehicle) {
	s.vehicle = &v
}

func (s SimpleParkingSpace) Vacate() {
	s.vehicle = nil
}

func (s SimpleParkingSpace) IsOccupied() (bool, IVehicle) {
	if s.vehicle == nil {
		return false, nil
	}
	return true, *s.vehicle
}

func (s SimpleParkingSpace) DoesSupport(vehicleType string) bool {
	_, found := s.supportedVehiclesTypes[vehicleType]
	return found
}

func (s SimpleParkingSpace) SupportedVehiclesTypes() map[string]bool {
	return s.supportedVehiclesTypes
}

func NewSimpleParkingSpace(id int, floor int, supportedTypes []string) *SimpleParkingSpace {
	vehicleTypes := make(map[string]bool)
	for _, vehicleType := range supportedTypes {
		vehicleTypes[vehicleType] = true
	}
	return &SimpleParkingSpace{
		id:                     id,
		floor:                  floor,
		supportedVehiclesTypes: vehicleTypes,
	}
}
