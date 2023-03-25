package main

import (
	"fmt"
	"strconv"
	"os"

	// "pcg/attributes"
	"pcg/character"
	"pcg/rolling"
	"pcg/dbService"
)

func main() {
	dbService.DbConnect()
	races, err := dbService.GetRaces()
	if err != nil {
		fmt.Errorf("GettingRacesBombed: %v", err)
		os.Exit(1)
	}

	var characterName string

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
	raceId := choiceNum-1
	fmt.Printf("You chose %d to build a %s\n", choiceNum, races[raceId].Name)

	raceAttributes, err := dbService.GetRaceAttributes(choiceNum - 1)
	//fmt.Printf("%+v\n", raceAttributes)

	var newChar character.Character
	newChar.Name = characterName
	newChar.RaceId = raceId
	newChar.Lvl = 3
	newChar.IQ = rolling.RollD6Bonus(raceAttributes.IQ, raceAttributes.IQBonus)
	newChar.ME = rolling.RollD6Bonus(raceAttributes.ME, raceAttributes.MEBonus)
	newChar.MA = rolling.RollD6Bonus(raceAttributes.ME, raceAttributes.MEBonus)
	newChar.PS = rolling.RollD6Bonus(raceAttributes.PS, raceAttributes.PSBonus)
	newChar.PP = rolling.RollD6Bonus(raceAttributes.PP, raceAttributes.PPBonus)
	newChar.PE = rolling.RollD6Bonus(raceAttributes.PE, raceAttributes.PEBonus)
	newChar.PB = rolling.RollD6Bonus(raceAttributes.PB, raceAttributes.PBBonus)
	newChar.Spd = rolling.RollD6Bonus(raceAttributes.Spd, raceAttributes.SpdBonus)
	newChar.HP = newChar.PE + rolling.RollD6Bonus(newChar.Lvl, 0)
	newChar.PPE = rolling.RollD6Bonus(raceAttributes.PPE, raceAttributes.PPEBonus)
	newChar.SpdDig = rolling.RollD6Bonus(raceAttributes.SpdDig, raceAttributes.SpdDigBonus)

	fmt.Printf("Name: %s \n", newChar.Name)
	fmt.Printf("Race: %s \n", races[raceId].Name)
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

	isDigger := raceId == 3 || raceId == 4
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
