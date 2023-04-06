package main

import (
	"fmt"
	"strings"
)

func main() {
	keepGoing := "Y"
	var characterId int64
	characterId = -1
	var raceId int
	raceId = -1
	fmt.Println("Welcome to Palladium Fantasy Character Builder!\n")
	for strings.ToLower(keepGoing) == "y" {
		startChoice := starter()
		if startChoice == "2" || strings.ToLower(startChoice) == "y" {
			raceId, characterId = builder()
			charId := occser(raceId, characterId)
			if charId != characterId {
				fmt.Println("poop")
			}
		} else {
			reviewer(startChoice)
		}
		fmt.Printf("\nWould you like to continue reviewing or building characters? (Y/n) : ")
		fmt.Scanln(&keepGoing)
	}
}
