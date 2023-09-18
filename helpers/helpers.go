package helpers

import (
	"fmt"

	"PALLADIUM_FCG/types"
)

func PrintCharacter(character types.Character) {
	fmt.Printf("Name: %s \n", character.Name)
	fmt.Printf("Race: %s \n", character.Race)
	fmt.Printf("Level: %d \n", character.Lvl)
	fmt.Printf("IQ: %d \n", character.IQ)
	fmt.Printf("ME: %d \n", character.ME)
	fmt.Printf("MA: %d \n", character.MA)
	fmt.Printf("PS: %d \n", character.PS)
	fmt.Printf("PP: %d \n", character.PP)
	fmt.Printf("PE: %d \n", character.PE)
	fmt.Printf("PB: %d \n", character.PB)
	fmt.Printf("Spd: %d \n", character.Spd)
	fmt.Printf("HP: %d \n", character.HP)
	fmt.Printf("PPE: %d \n", character.PPE)
	fmt.Printf("HF: %d \n", character.HF)

	isDigger := character.SpdDig != 0
	if isDigger {
		fmt.Printf("Spd Digging: %d \n", character.SpdDig)
	}

	fmt.Printf("OCC: %s \n", character.OccDesc)
}
