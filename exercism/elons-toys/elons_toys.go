package elon

import "fmt"

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	newBattery := car.battery - car.batteryDrain
	if newBattery >= 0 {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

// TODO: define the 'DisplayDistance() string' method
func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.speed*(100-car.battery)/car.batteryDrain)
}

// TODO: define the 'DisplayBattery() string' method
func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car *Car) CanFinish(trackDistance int) bool {
	return car.battery*car.speed/car.batteryDrain >= trackDistance
}