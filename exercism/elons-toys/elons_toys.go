package elon

// TODO: define the 'Drive()' method
func (car *Car) Drive(speed, batteryDrain int) {
	car.distance += speed
	car.battery -= batteryDrain
}

// TODO: define the 'DisplayDistance() string' method
func (car *Car) DisplayDistance() string {
	panic("")
	//return "5" //work more on this
}

// TODO: define the 'DisplayBattery() string' method
func (car *Car) DisplayBattery() string {
	panic("")
	//return fmt.Sprintf("Battery is at %d", car.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
