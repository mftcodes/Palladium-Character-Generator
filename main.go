package main

import (
	"fmt"

	"pcg/attributes"
	"pcg/character"
	"pcg/rolling"
)

func main() {
	myRaceAtt := attributes.BirthElf()
	// fmt.Printf("%+v\n", myRaceAtt)
	var myChar character.Character
	myChar.Name = "Pat"
	myChar.Race = myRaceAtt.Name
	myChar.Lvl = 1
	myChar.IQ = rolling.Roll(myRaceAtt.IQ, myRaceAtt.IQBonus)
	myChar.ME = rolling.Roll(myRaceAtt.ME, myRaceAtt.MEBonus)
	myChar.MA = rolling.Roll(myRaceAtt.ME, myRaceAtt.MEBonus)
	myChar.PS = rolling.Roll(myRaceAtt.PS, myRaceAtt.PSBonus)
	myChar.PP = rolling.Roll(myRaceAtt.PP, myRaceAtt.PPBonus)
	myChar.PE = rolling.Roll(myRaceAtt.PE, myRaceAtt.PEBonus)
	myChar.PB = rolling.Roll(myRaceAtt.PB, myRaceAtt.PBBonus)
	myChar.Spd = rolling.Roll(myRaceAtt.Spd, myRaceAtt.SpdBonus)
	myChar.HP = myChar.PE + rolling.Roll(myRaceAtt.HP, myRaceAtt.HPBonus)
	myChar.PPE = rolling.Roll(myRaceAtt.PPE, myRaceAtt.PPEBonus)

	if myRaceAtt.SpdDig != nil {
		myChar.SpdDig = rolling.Roll(myRaceAtt.SpdDig, myRaceAtt.SpdDigBonus)
	}

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

	if hasMoreAttrs {
		fmt.Printf("Spd Digging: %d \n", myChar.SpdDig)
	}
}
