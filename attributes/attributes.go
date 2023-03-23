package attributes

type Race struct {
	Name      string //
	IQ        string //  Intelligence Quotient
	IQBonus   int    //
	ME        string // Mental Endurance
	MEBonus   int    //
	MA        string // Mental Affinity
	MABonus   int    //
	PS        string // Physical Strength
	PSBonus   int    //
	PP        string // Physical Prowess
	PPBonus   int    //
	PE        string // Physical Endurance
	PEBonus   int    //
	PB        string // Physical Beauty
	PBBonus   int    //
	Spd       string // Speed
	SpdBonus  int    //
	HP        string // Hit  Points
	HPBonus   int    //
	PPE       string // Potential Psychic  Energy
	PPEBonus  int    //
	Alignment string //
}

// type Alignment struct {
//		Could probably do an  enume or dictionary  here?
// }

type Human struct {
	Race
}

func BirthHuman() Human {
	return Human{
		Race{
			"Human",
			"3D6", // IQ
			0,     // IQ Bonus
			"3D6", // ME
			0,     // ME Bonus
			"3D6", // MA
			0,     // MA Bonus
			"3D6", // PS
			0,     // PS Bonus
			"3D6", // PP
			0,     // PP Bonus
			"3D6", // PE
			0,     // PE Bonus
			"3D6", // PB
			0,     // PB Bonus
			"3D6", // Spd
			0,     // Spd Bonus
			"1D6", // HP
			0,     // HP Bonus
			"2D6", // PPE - Children up to age 18 is 5D6, mage/clergy look to OCC
			0,     // PPE Bonus
			"Any, usually lean toward good and selfish",
		},
	}
}

type Elf struct {
	Race
}

func BirthElf() Elf {
	return Elf{
		Race{
			"Elf",
			"3D6", // IQ
			1,     // IQ Bonus
			"3D6", // ME
			0,     // ME Bonus
			"2D6", // MA
			0,     // MA Bonus
			"3D6", // PS
			0,     // PS Bonus
			"4D6", // PP
			0,     // PP Bonus
			"3D6", // PE
			0,     // PE Bonus
			"5D6", // PB
			0,     // PB Bonus
			"3D6", // Spd
			0,     // Spd Bonus
			"1D6", // HP
			0,     // HP Bonus
			"2D6", // PPE - Children up to age 18 is 5D6, mage/clergy look to OCC
			0,     // PPE Bonus
			"Any, usually lean toward good and selfish",
		},
	}
}

type Dwarf struct {
	Race
	SpdDig      string // Speed Digging
	SpdDigBonus int    // Speed Digging Bonus
}

func BirthDwarf() Dwarf {
	return Dwarf{
		Race{
			"Dwarf",
			"3D6", // IQ
			0,     // IQ Bonus
			"3D6", // ME
			0,     // ME Bonus
			"2D6", // MA
			0,     // MA Bonus
			"4D6", // PS
			6,     // PS Bonus
			"3D6", // PP
			0,     // PP Bonus
			"4D6", // PE
			0,     // PE Bonus
			"2D6", // PB
			2,     // PB Bonus
			"2D6", // Spd
			0,     // Spd Bonus
			"1D6", // HP
			0,     // HP Bonus
			"2D6", // PPE - Children up to age 18 is 5D6, mage/clergy look to OCC
			0,     // PPE Bonus
			"Any, usually lean toward good and selfish",
		},
		"1D6",
		0,
	}
}
