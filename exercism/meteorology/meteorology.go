package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (tu TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[tu]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t Temperature) String() string {
	return fmt.Sprintf("%v %v", t.degree, t.unit.String())
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit

func (su SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[su]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s Speed) String() string {
	return fmt.Sprintf("%v %v", s.magnitude, s.unit.String())
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (md MeteorologyData) String() string {
	return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity", md.location, md.temperature.String(), md.windDirection, md.windSpeed.String(), md.humidity)
}
