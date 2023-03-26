package attributes

type Race struct {
	Id	 int
	Name string
}

type RaceAttributes struct { // provide majority of base stats
	Id			int
	RaceId      int //
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
