package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averagePrepTime int) int {
	if averagePrepTime <= 0 || len(layers) == 0 {
		return 12
	}
	return len(layers) * averagePrepTime

}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var grams int
	var liters float64

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			grams += 50
		}
		if layers[i] == "sauce" {
			liters += 0.2
		}
	}
	return grams, liters
}

// TODO: define the 'AddSecretIngredient()' function

// TODO: define the 'ScaleRecipe()' function
