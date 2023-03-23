package attributes

type Race struct { // provide majority of base stats
	Name        string //
	IQ          int    //  Intelligence Quotient
	IQBonus     int    //
	ME          int    // Mental Endurance
	MEBonus     int    //
	MA          int    // Mental Affinity
	MABonus     int    //
	PS          int    // Physical Strength
	PSBonus     int    //
	PP          int    // Physical Prowess
	PPBonus     int    //
	PE          int    // Physical Endurance
	PEBonus     int    //
	PB          int    // Physical Beauty
	PBBonus     int    //
	Spd         int    // Speed
	SpdBonus    int    //
	PPE         int    // Potential Psychic  Energy
	PPEBonus    int    //
	Alignment   string //
	SpdDig      int    // Speed Digging
	SpdDigBonus int    // Speed Digging Bonus
}

// type Alignment struct {
//		Could probably do an  enume or dictionary  here?
// }

func BuildHuman() Race {
	return Race{
		"Human",
		3, // IQ
		0, // IQ Bonus
		3, // ME
		0, // ME Bonus
		3, // MA
		0, // MA Bonus
		3, // PS
		0, // PS Bonus
		3, // PP
		0, // PP Bonus
		3, // PE
		0, // PE Bonus
		3, // PB
		0, // PB Bonus
		3, // Spd
		0, // Spd Bonus
		2, // PPE - Children up to age 18 is 5D6, mage/clergy look to OCC
		0, // PPE Bonus
		"Any, usually lean toward good and selfish", // Allignment
		0, // SpdDig
		0, // SpdDig Bonus
	}
}

func BuildElf() Race {
	return Race{
		"Elf",
		3, // IQ
		1, // IQ Bonus
		3, // ME
		0, // ME Bonus
		2, // MA
		0, // MA Bonus
		3, // PS
		0, // PS Bonus
		4, // PP
		0, // PP Bonus
		3, // PE
		0, // PE Bonus
		5, // PB
		0, // PB Bonus
		3, // Spd
		0, // Spd Bonus
		2, // PPE - Children up to age 18 is 5D6, mage/clergy look to OCC
		0, // PPE Bonus
		"Any, usually lean toward good and selfish", // Allignment
		0, // SpdDig
		0, // SpdDig Bonus
	}
}

func BuildDwarf() Race {
	return Race{
		"Dwarf",
		3, // IQ
		0, // IQ Bonus
		3, // ME
		0, // ME Bonus
		2, // MA
		0, // MA Bonus
		4, // PS
		6, // PS Bonus
		3, // PP
		0, // PP Bonus
		4, // PE
		0, // PE Bonus
		2, // PB
		2, // PB Bonus
		2, // Spd
		0, // Spd Bonus
		2, // PPE - Children up to age 18 is 5D6, mage/clergy look to OCC
		0, // PPE Bonus
		"Any, usually lean toward good and selfish", // Allignment
		1, // SpdDig
		0, // SpdDig Bonus
	}
}

func BuildGnome() Race {
	return Race{
		"Gnome",
		3, // IQ
		0, // IQ Bonus
		1, // ME
		6, // ME Bonus
		3, // MA
		4, // MA Bonus
		1, // PS
		4, // PS Bonus
		4, // PP
		0, // PP Bonus
		3, // PE
		6, // PE Bonus
		4, // PB
		0, // PB Bonus
		2, // Spd
		0, // Spd Bonus
		2, // PPE - Children up to age 18 is 5D6, mage/clergy look to OCC
		0, // PPE Bonus
		"Any, but most tend to be good or selfish; an evil gnome is a rarity", // Allignment
		1, // SpdDig
		0, // SpdDig Bonus
	}
}

func BuildTroglodyte() Race {
	return Race{
		"Troglodyte",
		2, // IQ
		0, // IQ Bonus
		2, // ME
		0, // ME Bonus
		3, // MA
		0, // MA Bonus
		4, // PS
		4, // PS Bonus
		3, // PP
		6, // PP Bonus
		3, // PE
		0, // PE Bonus
		2, // PB
		0, // PB Bonus
		6, // Spd
		0, // Spd Bonus
		2, // PPE - Children up to age 18 is 5D6, mage/clergy look to OCC
		0, // PPE Bonus
		"Any, but most tend to be good or unprincipled; a selfish or evil troglodyte is a rarity", // Allignment
		3, // SpdDig
		0, // SpdDig Bonus
	}
}

func BuildKobold() Race {
	return Race{
		"Kobold",
		3, // IQ
		0, // IQ Bonus
		2, // ME
		0, // ME Bonus
		3, // MA
		0, // MA Bonus
		3, // PS
		3, // PS Bonus
		3, // PP
		0, // PP Bonus
		3, // PE
		0, // PE Bonus
		1, // PB
		6, // PB Bonus
		3, // Spd
		0, // Spd Bonus
		4, // PPE - Children up to age 18 is 5D6, mage/clergy look to OCC
		0, // PPE Bonus
		"Typically anarchist or evil, but most player characters are likely to be unprincipled, anarchist, aberrant or even good", // Allignment
		1, // SpdDig
		0, // SpdDig Bonus
	}
}
