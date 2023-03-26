package main

import (
	"fmt"
	"os"
	"strconv"

	"pcg/character"
	"pcg/dbService"
	"pcg/rolling"
)

func main() {
	dbService.DbConnect()

	var characterName string
	d6 := 6
	d4 := 4

	races, err := dbService.GetRaces()
	if err != nil {
		fmt.Errorf("GettingRacesBombed: %v", err)
		os.Exit(1)
	}

	fmt.Println("Let's start by building a character! First please type in a name.")
	fmt.Scanln(&characterName)
	fmt.Printf("Great! You're new character will be named %s!\n", characterName)
	fmt.Println("Now please pick a race.")

	for _, race := range races {
		fmt.Printf("Press %d for %s\n", race.Id, race.Name)
	}

	var choice string
	fmt.Scanln(&choice)

	choiceNum, _ := strconv.Atoi(choice)
	isHuman := choiceNum == 1
	isHobGoblin := choiceNum == 8
	fmt.Printf("You chose %d to build a %s\n", choiceNum, races[choiceNum-1].Name)

	raceAttributes, err := dbService.GetRaceAttributes(choiceNum)
	//fmt.Printf("%+v\n", raceAttributes)

	fmt.Printf("What level do you want to start at? (Typically 1, 2, or 3)\n")
	var levelChoice string
	fmt.Scanln(&levelChoice)
	level, _ := strconv.Atoi(levelChoice)

	var newChar character.Character
	newChar.Name = characterName
	newChar.RaceId = choiceNum
	newChar.Lvl = level
	fmt.Printf("\nRolling for IQ with %dD6, with bonus of +%d\n", raceAttributes.IQ, raceAttributes.IQBonus)
	newChar.IQ = rolling.RollAttributes(isHuman, d6, raceAttributes.IQ, raceAttributes.IQBonus)

	fmt.Printf("\nRolling for ME with %dD6, with bonus of +%d\n", raceAttributes.ME, raceAttributes.MEBonus)
	newChar.ME = rolling.RollAttributes(isHuman, d6, raceAttributes.ME, raceAttributes.MEBonus)

	fmt.Printf("\nRolling for MA with %dD6, with bonus of +%d\n", raceAttributes.MA, raceAttributes.MABonus)
	newChar.MA = rolling.RollAttributes(isHuman, d6, raceAttributes.MA, raceAttributes.MABonus)

	fmt.Printf("\nRolling for PS with %dD6, with bonus of +%d\n", raceAttributes.PS, raceAttributes.PSBonus)
	newChar.PS = rolling.RollAttributes(isHuman, d6, raceAttributes.PS, raceAttributes.PSBonus)

	fmt.Printf("\nRolling for PP with %dD6, with bonus of +%d\n", raceAttributes.PP, raceAttributes.PPBonus)
	newChar.PP = rolling.RollAttributes(isHuman, d6, raceAttributes.PP, raceAttributes.PPBonus)

	fmt.Printf("\nRolling for PE with %dD6, with bonus of +%d\n", raceAttributes.PE, raceAttributes.PEBonus)
	newChar.PE = rolling.RollAttributes(isHuman, d6, raceAttributes.PE, raceAttributes.PEBonus)

	fmt.Printf("\nRolling for PB with %dD6, with bonus of +%d\n", raceAttributes.PB, raceAttributes.PBBonus)
	newChar.PB = rolling.RollAttributes(isHuman, d6, raceAttributes.PB, raceAttributes.PBBonus)

	fmt.Printf("\nRolling for Spd with %dD6, with bonus of +%d\n", raceAttributes.Spd, raceAttributes.SpdBonus)
	newChar.Spd = rolling.RollAttributes(isHuman, d6, raceAttributes.Spd, raceAttributes.SpdBonus)

	fmt.Printf("\nRolling for HP with %dD6, with bonus of +%d\n", newChar.Lvl, 0)
	newChar.HP = newChar.PE + rolling.RollAttributes(isHuman, d6, newChar.Lvl, 0)

	fmt.Printf("\nRolling for PPE with %dD6, with bonus of +%d\n", raceAttributes.PPE, raceAttributes.PPEBonus)
	newChar.PPE = rolling.RollAttributes(isHuman, d6, raceAttributes.PPE, raceAttributes.PPEBonus)
	
	if isHobGoblin {
		fmt.Printf("\nRolling for Spd Digging with %dD4, with bonus of +%d\n", raceAttributes.SpdDig, raceAttributes.SpdDigBonus)
		newChar.SpdDig = rolling.RollAttributes(isHuman, d4, raceAttributes.SpdDig, raceAttributes.SpdDigBonus)
	} else {
		fmt.Printf("\nRolling for Spd Digging with %dD6, with bonus of +%d\n", raceAttributes.SpdDig, raceAttributes.SpdDigBonus)
		newChar.SpdDig = rolling.RollAttributes(isHuman, d6, raceAttributes.SpdDig, raceAttributes.SpdDigBonus)
	}

	fmt.Printf("\n\nName: %s \n", newChar.Name)
	fmt.Printf("Race: %s \n", races[choiceNum-1].Name)
	fmt.Printf("Level: %d \n", newChar.Lvl)
	fmt.Printf("IQ: %d \n", newChar.IQ)
	fmt.Printf("ME: %d \n", newChar.ME)
	fmt.Printf("MA: %d \n", newChar.MA)
	fmt.Printf("PS: %d \n", newChar.PS)
	fmt.Printf("PP: %d \n", newChar.PP)
	fmt.Printf("PE: %d \n", newChar.PE)
	fmt.Printf("PB: %d \n", newChar.PB)
	fmt.Printf("Spd: %d \n", newChar.Spd)
	fmt.Printf("HP: %d \n", newChar.HP)
	fmt.Printf("PPE: %d \n", newChar.PPE)

	isDigger := choiceNum >= 3 && choiceNum <= 8
	if isDigger {
		fmt.Printf("Spd Digging: %d \n", newChar.SpdDig)
	}

	fmt.Println("Saving Character.")
	newCharId, err := dbService.SaveCharacter(newChar)
	if err != nil {
		fmt.Printf("PPE: %v \n", err)
		fmt.Errorf("ErrorSavingChar: %v", err)
		os.Exit(1)
	}

	fmt.Printf("New Character saved with id %d", newCharId)
}
