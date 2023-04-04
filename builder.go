package main

import (
	"fmt"
	"os"
	"strconv"
)

func builder() {
	dbConnect()

	var choice string
	d6 := 6
	d4 := 4

	races, err := getRaces()
	if err != nil {
		fmt.Errorf("GettingRacesBombed: %v", err)
		os.Exit(1)
	}

	var characterName string
	fmt.Println("Let's building a character! First please type in a name.")
	fmt.Scanln(&characterName)
	fmt.Printf("Great! You're new character will be named %s!\n", characterName)
	fmt.Println("Now please pick a race.")

	for _, race := range races {
		fmt.Printf("Press %d for %s\n", race.Id, race.Name)
	}
	fmt.Scanln(&choice)

	choiceNum, _ := strconv.Atoi(choice)
	isHuman := choiceNum == 1
	isHobGoblin := choiceNum == 8
	fmt.Printf("You chose %d to build a %s\n", choiceNum, races[choiceNum-1].Name)

	raceAttributes, err := getRaceAttributes(choiceNum)
	//fmt.Printf("%+v\n", raceAttributes)

	fmt.Printf("What level do you want to start at? (Typically 1, 2, or 3)\n")
	var levelChoice string
	fmt.Scanln(&levelChoice)
	level, _ := strconv.Atoi(levelChoice)

	var newChar character
	newChar.Name = characterName
	newChar.RaceId = choiceNum
	newChar.Race = races[choiceNum-1].Name
	newChar.Lvl = level
	fmt.Printf("\nRolling for IQ with %dD6, with bonus of +%d\n", raceAttributes.IQ, raceAttributes.IQBonus)
	newChar.IQ = rollAttributes(isHuman, d6, raceAttributes.IQ, raceAttributes.IQBonus)

	fmt.Printf("\nRolling for ME with %dD6, with bonus of +%d\n", raceAttributes.ME, raceAttributes.MEBonus)
	newChar.ME = rollAttributes(isHuman, d6, raceAttributes.ME, raceAttributes.MEBonus)

	fmt.Printf("\nRolling for MA with %dD6, with bonus of +%d\n", raceAttributes.MA, raceAttributes.MABonus)
	newChar.MA = rollAttributes(isHuman, d6, raceAttributes.MA, raceAttributes.MABonus)

	fmt.Printf("\nRolling for PS with %dD6, with bonus of +%d\n", raceAttributes.PS, raceAttributes.PSBonus)
	newChar.PS = rollAttributes(isHuman, d6, raceAttributes.PS, raceAttributes.PSBonus)

	fmt.Printf("\nRolling for PP with %dD6, with bonus of +%d\n", raceAttributes.PP, raceAttributes.PPBonus)
	newChar.PP = rollAttributes(isHuman, d6, raceAttributes.PP, raceAttributes.PPBonus)

	fmt.Printf("\nRolling for PE with %dD6, with bonus of +%d\n", raceAttributes.PE, raceAttributes.PEBonus)
	newChar.PE = rollAttributes(isHuman, d6, raceAttributes.PE, raceAttributes.PEBonus)

	fmt.Printf("\nRolling for PB with %dD6, with bonus of +%d\n", raceAttributes.PB, raceAttributes.PBBonus)
	newChar.PB = rollAttributes(isHuman, d6, raceAttributes.PB, raceAttributes.PBBonus)

	fmt.Printf("\nRolling for Spd with %dD6, with bonus of +%d\n", raceAttributes.Spd, raceAttributes.SpdBonus)
	newChar.Spd = rollAttributes(isHuman, d6, raceAttributes.Spd, raceAttributes.SpdBonus)

	fmt.Printf("\nRolling for HP with %dD6, with bonus of +%d\n", newChar.Lvl, 0)
	newChar.HP = newChar.PE + rollAttributes(isHuman, d6, newChar.Lvl, 0)

	fmt.Printf("\nRolling for PPE with %dD6, with bonus of +%d\n", raceAttributes.PPE, raceAttributes.PPEBonus)
	newChar.PPE = rollAttributes(isHuman, d6, raceAttributes.PPE, raceAttributes.PPEBonus)

	if isHobGoblin {
		fmt.Printf("\nRolling for Spd Digging with %dD4, with bonus of +%d\n", raceAttributes.SpdDig, raceAttributes.SpdDigBonus)
		newChar.SpdDig = rollAttributes(isHuman, d4, raceAttributes.SpdDig, raceAttributes.SpdDigBonus)
	} else {
		fmt.Printf("\nRolling for Spd Digging with %dD6, with bonus of +%d\n", raceAttributes.SpdDig, raceAttributes.SpdDigBonus)
		newChar.SpdDig = rollAttributes(isHuman, d6, raceAttributes.SpdDig, raceAttributes.SpdDigBonus)
	}

	fmt.Printf("\n\n")
	printCharacter(newChar)

	fmt.Println("Saving Character.")
	newCharId, err := saveCharacter(newChar)
	if err != nil {
		fmt.Printf("Save error: %v \n", err)
		fmt.Errorf("ErrorSavingChar: %v", err)
		os.Exit(1)
	}

	fmt.Printf("New Character saved with id %d\n", newCharId)
}
