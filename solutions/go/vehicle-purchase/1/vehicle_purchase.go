package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	vehichle := option1

	if option2 < option1 {
		vehichle = option2
	}

	return fmt.Sprintf("%s is clearly the better choice.", vehichle)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	switch {
	case age >= 10:
		return originalPrice * 0.5
	case age < 10 && age >= 3:
		return originalPrice * 0.7
	case age <= 3:
		return originalPrice * 0.8
	default:
		return originalPrice
	}
}
