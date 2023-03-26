package rolling

import (
	"fmt"
	"math/rand"
	"time"
)

func Random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func RollAttributes(isHuman bool, sides, numDice, bonus int) int { // isHuman, sides, number, bonus
	total := 0

	for i := 1; i <= numDice; i++ {
		roll := Random(1, sides+1)
		fmt.Printf("roll %d for %d.\n", i, roll)
		total += roll
	}

	if isHuman && numDice == 3 && total >= 16 {
		total += exceptionalRoll(sides, isHuman)
	} else if numDice == 2 && total == 12 {
		total += exceptionalRoll(sides, isHuman)
	}

	return total + bonus
}

func exceptionalRoll(sides int, isHuman bool) int {
	fmt.Printf("Exceptional roll! Rolling 1D%d again.\n", sides)
	roll := Random(1, sides+1)
	if isHuman && roll == 6 {
		fmt.Println("The universe is pleased with you! You rolled a 6, rollin again!")
		roll += Random(1, sides+1)
		fmt.Printf("Second roll was %d, adding a additional %d to your attribute.\n", roll-6, roll)
		return roll
	}
	fmt.Printf("Rolled %d will be added to your attribute.\n", roll)
	return roll
}

// This is cool, but overkill. Do not need this much abstraction and flexibility.
// func RollFlex(diceAttrs ...int) int { // isHuman, sides, number, bonus
// 	sides := diceAttrs[1]
// 	numDice := diceAttrs[2]
// 	bonus := 0

// 	numSides := sides+1
// 	total := 0

// 	for i := 1; i<= numDice; i++ {
// 		roll := Random(1, numSides)
// 		fmt.Printf("roll %d for %d.\n", i, roll)
// 		total += roll
// 	}

// 	if len(diceAttrs) == 3 {
// 		bonus = diceAttrs[2]
// 		fmt.Printf("bonus of %d will be added\n", bonus)
// 	}

// 	return total+bonus
// }
