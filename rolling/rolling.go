package rolling

import (
	"math/rand"
)

func RollD6Bonus(numDice int, bonus int) int {
	total := RollD6(numDice) + bonus
	return total
}

func RollD6(numDice int) int {
	total := 0

	for i := 0; i < numDice; i++ {
		total += Random (1, 7)
	}
	
	switch numDice {
	case 1:
		if total == 6 {
			total += Random(1, 7)
		}
	case 2:
		if total >= 11 {
			total += Random(1, 7)
		}
	case 3:
		if total >= 16 {
			total += Random(1, 7)
		}
	case 4:
		if total >= 21 {
			total += Random(1, 7)
		}
	case 5:
		if total >= 21 {
			total += Random(1, 7)
		}
	default:
	}

	return total;
}

func Random(min int, max int) int {
	return rand.Intn(max-min) + min
}
