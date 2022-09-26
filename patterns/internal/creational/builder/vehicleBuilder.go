package builder

// The VehicleDirector director variable is the one in charge of accepting the builders.
// It has a Construct method that will use the builder that is stored in Vehicle,
// and will reproduce the required steps.
type VehicleDirector struct {
	builder BuilderProcess
}

func (vd *VehicleDirector) SetBuilder(b BuilderProcess) {
	vd.builder = b
}

func (vd *VehicleDirector) Construct() {
	vd.builder.SetSeats().SetColor().SetType().SetWheels()
}

// This interface defines the steps that are necessary to build a vehicle.
type BuilderProcess interface {
	SetWheels() BuilderProcess
	SetSeats() BuilderProcess
	SetColor() BuilderProcess
	SetType() BuilderProcess
	Get() *VehicleProduct
}

type VehicleProduct struct {
	Type   string
	Wheels byte
	Seats  uint16
	Color  string
}
