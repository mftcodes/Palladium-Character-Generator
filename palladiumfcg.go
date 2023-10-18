package main

import (
	"fmt"
	"strings"

	"pfcg/builder"
	"pfcg/occ"
	"pfcg/prompts"
	"pfcg/reviewer"
	"pfcg/types"
)

func main() {
	keepGoing := "Y"
	choice := "Y"
	var characterId int64
	characterId = -1
	var raceId int
	raceId = -1
	var character types.Character

	fmt.Println("Welcome to Palladium Fantasy Character Builder!")
	fmt.Println("Just a simple application to help you quickly setup a new character.")
	fmt.Println()

	for strings.ToLower(keepGoing) == "y" {
		startChoice := prompts.Starter()
		if startChoice == "2" || strings.ToLower(startChoice) == "y" {
			raceId, characterId = builder.Builder()
			charId := occ.Occser(raceId, characterId)
			if charId != characterId {
				fmt.Println("Oh schnikes, something's amiss")
			}
		} else {
			character = reviewer.Reviewer(startChoice)
			characterId = int64(character.Id)
			raceId = character.RaceId

			if character.OccId == 1 {
				fmt.Printf("\nLooks like you have not yet chosen an OCC, would you like to do that now? (Y/n) ")
				fmt.Scanln(&choice)

				if strings.ToLower(choice) == "y" {
					charId := occ.Occser(raceId, characterId)
					if charId != characterId {
						fmt.Println("Oh schnikes, something's amiss")
					}
				}
			}
		}
		fmt.Printf("\nWould you like to continue reviewing or building characters? (Y/n) : ")
		fmt.Scanln(&keepGoing)
	}
}
