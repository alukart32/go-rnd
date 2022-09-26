package builder

import (
	"testing"
)

func TestVehicleCarBuilder(t *testing.T) {
	vehicleBuilder := VehicleDirector{}

	carBuilder := &CarBuilder{}
	carBuilder.SetProduct(&VehicleProduct{})
	vehicleBuilder.SetBuilder(carBuilder)
	vehicleBuilder.Construct()

	car := carBuilder.Get()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}

	if car.Type != "car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Type)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}

	motorbikeBuilder := &MotorbikeBuilder{}
	motorbikeBuilder.SetProduct(&VehicleProduct{})
	vehicleBuilder.SetBuilder(motorbikeBuilder)
	vehicleBuilder.Construct()

	motorbike := motorbikeBuilder.Get()

	if motorbike.Wheels != 2 {
		t.Errorf("Wheels on a car must be 2 and they were %d\n", motorbike.Wheels)
	}

	if motorbike.Type != "motorbike" {
		t.Errorf("Structure on a car must be 'Motorbike' and was %s\n", motorbike.Type)
	}

	if motorbike.Seats != 1 {
		t.Errorf("Seats on a car must be 1 and they were %d\n", motorbike.Seats)
	}
}
