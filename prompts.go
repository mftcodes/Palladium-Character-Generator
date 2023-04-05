package main

import (
	"fmt"
	"os"
)

func starter() string {
	dbConnect()

	var choice string

	numberCharactersSaved, err := getCharacterCount()
	if err != nil {
		fmt.Printf("error getting character count: %v \n", err)
		fmt.Errorf("characterCount: %v", err)
		os.Exit(1)
	}

	if numberCharactersSaved > 0 {
		fmt.Printf("You have %d characters saved.\n", numberCharactersSaved)
		fmt.Printf("Type a name to search for a character by name.\n")
		fmt.Printf("Press 1 to list all characters and select one from the list.\n")
		fmt.Printf("Press 2 to create a new character.\n")
	} else {
		fmt.Printf("No characters saved at yet, so let's go to the builder! Press 'Y' when ready: ")
	}
	fmt.Scanln(&choice)

	return choice
}
