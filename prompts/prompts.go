package prompts

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"pfcg/dbservice"
	"pfcg/types"
	"golang.org/x/exp/slices"
)

func Starter() string {
	dbservice.Connect()

	choice := "y"

	numberCharactersSaved, err := dbservice.GetCharacterCount()
	if err != nil {
		fmt.Printf("error getting character count: %v \n", err)
		fmt.Errorf("characterCount: %v", err)
		os.Exit(1)
	}

	if numberCharactersSaved > 0 {
		fmt.Printf("You have %d characters saved.\n", numberCharactersSaved)
		fmt.Printf("You have options at this point, please respond accordingly:\n")
		fmt.Printf("You can search for a known saved character by typing their name,\n")
		fmt.Printf("OR press 1 to list all saved characters and select one from the list,\n")
		fmt.Printf("OR press 2 to create a new character.\n")
		fmt.Printf("Type your answer here then press 'Enter': ")
	} else {
		fmt.Printf("No characters saved at yet, would you like to build one now? (Y/n) ")
	}
	fmt.Scanln(&choice)
	fmt.Println()

	return choice
}

func SetCharacterName() string {
	var first, middle, last string

	fmt.Printf("Let's building a character!\nFirst we need a name, limited to first, last, or first, middle, last at this time.\n")
	fmt.Printf("Type the name of your character: ")
	fmt.Scanln(&first, &middle, &last)
	characterName := strings.TrimSpace(fmt.Sprintf("%s %s %s", first, middle, last))
	fmt.Printf("Great! You're new character will be named %s\n", characterName)
	return characterName
}

func SetCharacterRace() (bool, bool, int, string) {
	choice := "11" //defaults to troll ( ͡ ͜ʖ ͡ )
	races, err := dbservice.GetRaces()
	if err != nil {
		fmt.Errorf("GettingRacesBombed: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Now we need to choose a character race.\nBelow are your listed racial options:\n")

SetRace:
	for _, race := range races {
		fmt.Printf("Choose %d for %s\n", race.Id, race.Desc)
	}
	fmt.Printf("Type the appropriate racial Id and press 'Enter' to set your race: ")
	fmt.Scanln(&choice)

	choiceNum, _ := strconv.Atoi(choice)
	isHuman := choiceNum == 1
	isHobGoblin := choiceNum == 8
	raceIndex := slices.IndexFunc(races, func(r types.Race) bool { return r.Id == choiceNum })
	raceDesc := races[raceIndex].Desc

	fmt.Printf("You chose %d to build a %s\n", choiceNum, raceDesc)
	choice = "y"
	fmt.Printf("If this is correct? (Y/n) ")
	fmt.Scanln(&choice)

	if choice != "y" {
		goto SetRace
	}

	return isHuman, isHobGoblin, choiceNum, raceDesc
}
