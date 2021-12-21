package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		avgPrepTime = 2
	}

	return len(layers) * avgPrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	n := 0
	s := 0.0
	for _, val := range layers {
		if val == "noodles" {
			n += 50
		} else if val == "sauce" {
			s += 0.2
		}
	}
	return n, s
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) []string {
	finalList := myList
	finalList = append(finalList, friendList[len(friendList)-1])
	return finalList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numPortions int) []float64 {
	newQuantities := make([]float64, len(quantities))
	for i, val := range quantities {
		newQuantities[i] = val * float64(numPortions) / 2
	}
	return newQuantities
}
