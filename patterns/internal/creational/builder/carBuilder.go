package builder

type CarBuilder struct {
	v *VehicleProduct
}

func (c *CarBuilder) SetProduct(p *VehicleProduct) {
	c.v = p
}

func (c *CarBuilder) SetType() BuilderProcess {
	c.v.Type = "car"
	return c
}

func (c *CarBuilder) SetWheels() BuilderProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuilderProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetColor() BuilderProcess {
	c.v.Color = "red"
	return c
}

func (c *CarBuilder) Get() *VehicleProduct {
	return c.v
}
