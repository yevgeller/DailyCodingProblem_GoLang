package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averagePrepTime int) int {
	if averagePrepTime <= 0 {
		averagePrepTime = 2
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
func AddSecretIngredient(friends, own []string) {
	own[len(own)-1] = friends[len(friends)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var result []float64
	// if len(quantities) == 0 {

	// 	return result
	// }
	for i := 0; i < len(quantities); i++ {
		result = append(result, quantities[i]*(float64(portions)/2))
	}
	return result
}
