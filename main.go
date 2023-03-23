package main

import (
	"fmt"

	"pcg/attributes"
	"pcg/character"
	"pcg/rolling"
)

func main() {
	myRaceAtt := attributes.BuildHuman()
	// fmt.Printf("%+v\n", myRaceAtt)
	var myChar character.Character
	myChar.Name = "Pat"
	myChar.Race = myRaceAtt.Name
	myChar.Lvl = 3
	myChar.IQ = rolling.RollD6Bonus(myRaceAtt.IQ, myRaceAtt.IQBonus)
	myChar.ME = rolling.RollD6Bonus(myRaceAtt.ME, myRaceAtt.MEBonus)
	myChar.MA = rolling.RollD6Bonus(myRaceAtt.ME, myRaceAtt.MEBonus)
	myChar.PS = rolling.RollD6Bonus(myRaceAtt.PS, myRaceAtt.PSBonus)
	myChar.PP = rolling.RollD6Bonus(myRaceAtt.PP, myRaceAtt.PPBonus)
	myChar.PE = rolling.RollD6Bonus(myRaceAtt.PE, myRaceAtt.PEBonus)
	myChar.PB = rolling.RollD6Bonus(myRaceAtt.PB, myRaceAtt.PBBonus)
	myChar.Spd = rolling.RollD6Bonus(myRaceAtt.Spd, myRaceAtt.SpdBonus)
	myChar.HP = myChar.PE + rolling.RollD6Bonus(myChar.Lvl, 0)
	myChar.PPE = rolling.RollD6Bonus(myRaceAtt.PPE, myRaceAtt.PPEBonus)
	myChar.SpdDig = rolling.RollD6Bonus(myRaceAtt.SpdDig, myRaceAtt.SpdDigBonus)

	fmt.Printf("Name: %s \n", myChar.Name)
	fmt.Printf("Race: %s \n", myChar.Race)
	fmt.Printf("Level: %d \n", myChar.Lvl)
	fmt.Printf("IQ: %d \n", myChar.IQ)
	fmt.Printf("ME: %d \n", myChar.ME)
	fmt.Printf("MA: %d \n", myChar.MA)
	fmt.Printf("PS: %d \n", myChar.PS)
	fmt.Printf("PP: %d \n", myChar.PP)
	fmt.Printf("PE: %d \n", myChar.PE)
	fmt.Printf("PB: %d \n", myChar.PB)
	fmt.Printf("Spd: %d \n", myChar.Spd)
	fmt.Printf("HP: %d \n", myChar.HP)
	fmt.Printf("PPE: %d \n", myChar.PPE)

	isDigger := myRaceAtt.Name == "Dwarf" || myRaceAtt.Name == "Gnome"
	if isDigger {
		fmt.Printf("Spd Digging: %d \n", myChar.SpdDig)
	}
}
