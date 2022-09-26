package builder

type MotorbikeBuilder struct {
	v *VehicleProduct
}

func (c *MotorbikeBuilder) SetProduct(p *VehicleProduct) {
	c.v = p
}

func (m *MotorbikeBuilder) SetType() BuilderProcess {
	m.v.Type = "motorbike"
	return m
}

func (m *MotorbikeBuilder) SetWheels() BuilderProcess {
	m.v.Wheels = 2
	return m
}

func (m *MotorbikeBuilder) SetSeats() BuilderProcess {
	m.v.Seats = 1
	return m
}

func (m *MotorbikeBuilder) SetColor() BuilderProcess {
	m.v.Color = "black"
	return m
}

func (m *MotorbikeBuilder) Get() *VehicleProduct {
	return m.v
}
