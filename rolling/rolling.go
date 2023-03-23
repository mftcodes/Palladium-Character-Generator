package rolling

import (
	"fmt"
	"math/rand"
)

func Roll(dice string, bonus int) int {
	total := 0;

	switch dice {
		case "1D6":
			total = Random(1,7)
			if (total == 6) {
				total += Random(1,7)
			}

		case "2D6":
			total = Random(1,7) + Random(1,7)
			if (total >= 11) {
				total += Random(1,7)
			}

		case "3D6":
			total = Random(1,7) + Random(1,7) + Random(1,7)
			if (total >= 16) {
				total += Random(1,7)
			}

		case "4D6":
			total = Random(1,7) + Random(1,7) + Random(1,7) + Random(1,7)
			if (total >= 21) {
				total += Random(1,7)
			}

		default:
			fmt.Println("Whoops, bad roll input.")
	}

	total += bonus
	return total
}

func Random(min int, max int) int {
	return rand.Intn(max-min) + min
}