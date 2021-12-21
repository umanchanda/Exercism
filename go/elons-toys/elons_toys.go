package elon

import "strconv"

// TODO: define the 'Drive()' method
func (c *Car) Drive() *Car {
	if c.battery < c.batteryDrain {
		return c
	}
	c.distance += c.speed
	c.battery -= c.batteryDrain
	return c
}

// TODO: define the 'DisplayDistance() string' method
func (c Car) DisplayDistance() string {
	return "Driven " + strconv.Itoa(c.distance) + " meters"
}

// TODO: define the 'DisplayBattery() string' method
func (c Car) DisplayBattery() string {
	return "Battery at " + strconv.Itoa(c.battery) + "%"
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c *Car) CanFinish(trackDistance int) bool {
	for c.battery > 1 {
		c.Drive()
	}
	return c.distance >= trackDistance
}
