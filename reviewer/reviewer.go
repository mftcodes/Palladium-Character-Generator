package reviewer

import (
	"fmt"
	"os"
	"strconv"

	"pfcg/dbservice"
	"pfcg/helpers"
	"pfcg/types"
)

func Reviewer(choice string) (character types.Character) {
	dbservice.Connect()
	var characters []types.CharacterShort
	var err error

	if choice == "1" {
		characters, err = dbservice.GetCharactersShort()
		if err != nil {
			fmt.Printf("error getting character names: %v \n", err)
			fmt.Errorf("characterNames: %v", err)
			os.Exit(1)
		}
		fmt.Printf("Here is your list of characters:\n")
		for _, _character := range characters {
			fmt.Printf("Press %d to see stats the %s named %s\n", _character.Id, _character.Race, _character.Name)
		}
		fmt.Scanln(&choice)
		charId, _ := strconv.Atoi(choice)
		character, err = dbservice.GetCharacterById(charId)
		if err != nil {
			fmt.Printf("Character of ID %s was not found.", choice)
			fmt.Printf("error getting character by id: %v \n", err)
			fmt.Errorf("characterById: %v", err)
			// os.Exit(1)
		} else {
			fmt.Printf("\nHere are your characters attributes:\n")
			helpers.PrintCharacter(character)
		}
	} else {
		character, err = dbservice.GetCharacterByName(choice)
		if err != nil {
			fmt.Printf("%s was not found.", choice)
			fmt.Printf("error getting character by name: %v \n", err)
			fmt.Errorf("characterByName: %v", err)
			// os.Exit(1)
		} else {
			fmt.Printf("\nHere are your characters attributes:\n")
			helpers.PrintCharacter(character)
		}
	}

	return character
}
