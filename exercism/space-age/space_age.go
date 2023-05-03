package space

import ("strings")

type Planet struct {
	name string
}

func Age(seconds float64, planet Planet) float64 {
	//panic("Please implement the Age function")

	earthAge := seconds / 31557600

switch(strings.ToLower(planet.name))

	if strings.ToLower(planet.name) == "earth" {
return earthAge
	} 
}
