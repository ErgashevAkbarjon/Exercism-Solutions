package lasagna

func PreparationTime(layers []string, averagePrepTime int) int {
	if averagePrepTime == 0 {
		averagePrepTime = 2
	}

	return len(layers) * averagePrepTime
}

func Quantities(layers []string) (int, float64) {
	var noodlesLayers int
	var sauceLayers float64

	for _, layer := range layers {
		if layer == "noodles" {
			noodlesLayers++
		}
		if layer == "sauce" {
			sauceLayers++
		}
	}

	return noodlesLayers * 50, sauceLayers * 0.2
}

func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
	scaledAmounts := make([]float64, len(amounts))

	for i, amount := range amounts {
		scaledAmounts[i] = (amount / 2) * float64(portions)
	}

	return scaledAmounts
}
