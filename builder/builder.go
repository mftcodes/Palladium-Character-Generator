package builder

import (
	"fmt"
	"os"
	"strconv"

	"PALLADIUM_FCG/dbservice"
	"PALLADIUM_FCG/helpers"
	"PALLADIUM_FCG/prompts"
	"PALLADIUM_FCG/roller"
	"PALLADIUM_FCG/types"
)

func Builder() (raceId int, characterId int64) {
	dbservice.Connect()

	d6 := 6
	d4 := 4

	characterName := prompts.SetCharacterName()
	fmt.Printf("Great! You're new character will be named %s!\n", characterName)

	isHuman, isHobGoblin, raceId, raceDesc := prompts.SetCharacterRace()
	fmt.Printf("\n******************************\n\nRaceId = %d\n\n", raceId)

	raceAttr, err := dbservice.GetRaceAttributes(raceId)
	fmt.Printf("%+v\n\n", raceAttr)

	fmt.Printf("What level do you want to start at? (Typically 1, 2, or 3)\n")
	var levelChoice string
	fmt.Scanln(&levelChoice)
	level, _ := strconv.Atoi(levelChoice)

	var newChar types.Character
	newChar.Name = characterName
	newChar.RaceId = raceId
	newChar.Race = raceDesc
	newChar.Lvl = level
	fmt.Printf("\nRolling for IQ with %dD6, with bonus of +%d\n", raceAttr.IQ, raceAttr.IQBonus)
	newChar.IQ = roller.RollAttributes(isHuman, d6, raceAttr.IQ, raceAttr.IQBonus)

	fmt.Printf("\nRolling for ME with %dD6, with bonus of +%d\n", raceAttr.ME, raceAttr.MEBonus)
	newChar.ME = roller.RollAttributes(isHuman, d6, raceAttr.ME, raceAttr.MEBonus)

	fmt.Printf("\nRolling for MA with %dD6, with bonus of +%d\n", raceAttr.MA, raceAttr.MABonus)
	newChar.MA = roller.RollAttributes(isHuman, d6, raceAttr.MA, raceAttr.MABonus)

	fmt.Printf("\nRolling for PS with %dD6, with bonus of +%d\n", raceAttr.PS, raceAttr.PSBonus)
	newChar.PS = roller.RollAttributes(isHuman, d6, raceAttr.PS, raceAttr.PSBonus)

	fmt.Printf("\nRolling for PP with %dD6, with bonus of +%d\n", raceAttr.PP, raceAttr.PPBonus)
	newChar.PP = roller.RollAttributes(isHuman, d6, raceAttr.PP, raceAttr.PPBonus)

	fmt.Printf("\nRolling for PE with %dD6, with bonus of +%d\n", raceAttr.PE, raceAttr.PEBonus)
	newChar.PE = roller.RollAttributes(isHuman, d6, raceAttr.PE, raceAttr.PEBonus)

	fmt.Printf("\nRolling for PB with %dD6, with bonus of +%d\n", raceAttr.PB, raceAttr.PBBonus)
	newChar.PB = roller.RollAttributes(isHuman, d6, raceAttr.PB, raceAttr.PBBonus)

	fmt.Printf("\nRolling for Spd with %dD6, with bonus of +%d\n", raceAttr.Spd, raceAttr.SpdBonus)
	newChar.Spd = roller.RollAttributes(isHuman, d6, raceAttr.Spd, raceAttr.SpdBonus)

	fmt.Printf("\nRolling for HP with %dD6, with bonus of +%d\n", newChar.Lvl, 0)
	newChar.HP = newChar.PE + roller.RollAttributes(isHuman, d6, newChar.Lvl, 0)

	fmt.Printf("\nRolling for PPE with %dD6, with bonus of +%d\n", raceAttr.PPE, raceAttr.PPEBonus)
	newChar.PPE = roller.RollAttributes(isHuman, d6, raceAttr.PPE, raceAttr.PPEBonus)

	newChar.HF = raceAttr.HF

	if isHobGoblin {
		fmt.Printf("\nRolling for Spd Digging with %dD4, with bonus of +%d\n", raceAttr.SpdDig, raceAttr.SpdDigBonus)
		newChar.SpdDig = roller.RollAttributes(isHuman, d4, raceAttr.SpdDig, raceAttr.SpdDigBonus)
	} else {
		fmt.Printf("\nRolling for Spd Digging with %dD6, with bonus of +%d\n", raceAttr.SpdDig, raceAttr.SpdDigBonus)
		newChar.SpdDig = roller.RollAttributes(isHuman, d6, raceAttr.SpdDig, raceAttr.SpdDigBonus)
	}

	fmt.Printf("\n\n")
	helpers.PrintCharacter(newChar)

	fmt.Println("Saving Character.")
	characterId, err = dbservice.SaveCharacter(newChar)
	if err != nil {
		fmt.Printf("Save error: %v \n", err)
		fmt.Errorf("ErrorSavingChar: %v", err)
		os.Exit(1)
	}

	fmt.Printf("New Character saved with id %d\n", characterId)
	return raceId, characterId
}
