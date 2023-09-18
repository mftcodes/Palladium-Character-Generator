package types

type Race struct {
	Id   int
	Desc string
}

type RaceAttr struct { // provide majority of base stats
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
	HF          int    // Horror Factor
	Alignment   string //
	SpdDig      int    // Speed Digging
	SpdDigBonus int    // Speed Digging Bonus
}

type Character struct {
	Id      int
	Name    string
	RaceId  int
	Race    string
	Lvl     int
	IQ      int
	ME      int
	MA      int
	PS      int
	PP      int
	PE      int
	PB      int
	Spd     int
	HP      int
	PPE     int
	HF      int // Horror Factor
	SpdDig  int
	OccId   int
	OccDesc string
}

type CharacterShort struct {
	Id   int
	Name string
	Race string
}

type SkillCategory struct {
	Id   int
	Desc string
}

type Skill struct {
	Id              int
	Desc            string
	SkillCategoryId int
}

type OccType struct {
	Id   int
	Desc string
}

type Occ struct {
	Id   int
	Type string
	Desc string
}

type RaceOcc struct {
	OccId  int
	RaceId int
}

type NaturalAbility struct {
	Id   int
	Desc int
}

type RaceNaturalAbility struct {
	NaturalAbiltyId int
	RaceId          int
	BonusInitial    int
	BonusPerLevel   int
	Value           int
	Measurement     string
	Note            string
}
