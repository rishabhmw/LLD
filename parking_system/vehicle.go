package parking_system

type IVehicle interface {
	GetType() string
	GetID() string
}

const (
	VehicleTypeCar   = "car"
	VehicleTypeTruck = "truck"
	VehicleTypeBike  = "bike"
)
