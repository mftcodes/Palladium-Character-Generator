package main

import (
	"fmt"
	"strings"
)

func main() {
	keepGoing := "Y"
	choice := "Y"
	var characterId int64
	characterId = -1
	var raceId int
	raceId = -1
	var character character

	fmt.Println("Welcome to Palladium Fantasy Character Builder!")
	fmt.Println("Just a simple application to help you quickly set-up a new character.")
	fmt.Println()

	for strings.ToLower(keepGoing) == "y" {
		startChoice := starter()
		if startChoice == "2" || strings.ToLower(startChoice) == "y" {
			raceId, characterId = builder()
			charId := occser(raceId, characterId)
			if charId != characterId {
				fmt.Println("Oh schnikes, something's amiss")
			}
		} else {
			character = reviewer(startChoice)
			characterId = int64(character.Id)
			raceId = character.RaceId

			if character.OccId == 1 {
				fmt.Printf("\nLooks like you have not yet chosen an OCC, would you like to do that now? (Y/n) ")
				fmt.Scanln(&choice)

				if strings.ToLower(choice) == "y" {
					charId := occser(raceId, characterId)
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
