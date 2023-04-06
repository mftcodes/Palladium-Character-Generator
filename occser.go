package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func occser(raceId  int, characterId int64) (int64) {
	var choice string
	var characterIdUpdated int64

	fmt.Println("Now that the base character is set up, let's see what OCC/PCC is available to choose.")
	
	occs, err := getOccsByRace(raceId)
	if err != nil {
		fmt.Printf("error getting OCCs by Race: %v \n", err)
		fmt.Errorf("occsByRace: %v", err)
		os.Exit(1)
	}

	CHOOSE:
	for _, occ := range occs {
		fmt.Printf("Press %d to see stats for %s: %s\n", occ.Id, occ.Type, occ.Desc)
	}
	fmt.Scanln(&choice)
	occId, _ := strconv.Atoi(choice)
	occ := occs[occId]
	fmt.Printf(`You chose %d and an OCC of  %s: %s will be assigned to your character. Is this correct?  (Y/n): `, occId, occ.Type, occ.Desc)
	fmt.Scanln(&choice)
	if strings.ToLower(choice) == "y" {
		characterIdUpdated, err = saveOcc(characterId, occId)
		if err != nil {
			fmt.Printf("error saving occ: %v \n", err)
			fmt.Errorf("occSave: %v", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("\nLet's try that again.\n")
		goto CHOOSE
	}
	return characterIdUpdated
}