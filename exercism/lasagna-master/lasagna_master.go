package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averagePrepTime int) int {
	if averagePrepTime <= 0 || len(layers) == 0 {
		return 12
	}
	return len(layers) * averagePrepTime

}

// TODO: define the 'Quantities()' function

// TODO: define the 'AddSecretIngredient()' function

// TODO: define the 'ScaleRecipe()' function
