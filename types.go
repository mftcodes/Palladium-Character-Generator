package main

type race struct {
	Id   int
	Desc string
}

type raceAttributes struct { // provide majority of base stats
	Id          int
	RaceId      int    //
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
	HF			int	   // Horror Factor
	Alignment   string //
	SpdDig      int    // Speed Digging
	SpdDigBonus int    // Speed Digging Bonus
}

type character struct {
	Id     int
	Name   string
	RaceId int
	Race   string
	Lvl    int
	IQ     int
	ME     int
	MA     int
	PS     int
	PP     int
	PE     int
	PB     int
	Spd    int
	HP     int
	PPE    int
	HF	   int	   // Horror Factor
	SpdDig int
}

type characterShort struct {
	Id   int
	Name string
}

