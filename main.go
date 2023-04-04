package main

import (
	"fmt"
	"strings"
)

func main() {
	keepGoing := "Y"

	fmt.Println("Welcome to Palladium Fantasy Character Builder!\n")
	for strings.ToLower(keepGoing) == "y" {
		startChoice := starter()
		if startChoice == "2" {
			builder()
		} else {
			reviewer(startChoice)
		}
		fmt.Printf("\nWould you like to continue reviewing or building characters? (Y/n) : ")
		fmt.Scanln(&keepGoing)
	}
}
