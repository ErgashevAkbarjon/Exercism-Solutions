package raindrops

import "strconv"

func Convert(number int) string {
	sound := ""
	sounds := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	for n, v := range sounds {
		if number%n == 0 {
			sound += v
		}
	}

	if sound == "" {
		sound = strconv.Itoa(number)
	}

	return sound
}
