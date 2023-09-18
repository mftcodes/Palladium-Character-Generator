package dbservice

import (
	"database/sql"
	"fmt"
	"log"

	"pfcg/types"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Connect
func Connect() {
	// fmt.Println("Connecting to Db")
	cfg := mysql.Config{
		User:   "root",
		Passwd: "pcg0superfun",
		Net:    "tcp",
		Addr:   "localhost:3308",
		DBName: "palladium",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	// fmt.Println("Connected and ready to begin!\n")
}

// Race table functions
func GetRaces() (races []types.Race, err error) {
	query := `SELECT * FROM Race`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("racesSelect: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var race types.Race
		if err := rows.Scan(&race.Id, &race.Desc); err != nil {
			return nil, fmt.Errorf("raceScan: %v", err)
		}
		races = append(races, race)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("racesFinal: %v", err)
	}
	return races, nil
}

func GetRaceByName(name string) (race types.Race, err error) {
	err = db.Ping()
	if err != nil {
		return types.Race{}, err
	}

	query := fmt.Sprintf(`SELECT * FROM Race r WHERE r.Desc = %s`, name)

	err = db.QueryRow(query).Scan(&race.Id, &race.Desc)
	if err != nil {
		return race, fmt.Errorf("raceSelectByName: %v", err)
	}

	return race, nil
}

func GetRaceById(id int) (race types.Race, err error) {
	err = db.Ping()
	if err != nil {
		return types.Race{}, err
	}

	query := fmt.Sprintf(`SELECT * FROM Race r WHERE r.Id = %d`, id)

	err = db.QueryRow(query).Scan(&race.Id, &race.Desc)
	if err != nil {
		return race, fmt.Errorf("raceSelectByName: %v", err)
	}

	return race, nil
}

func GetRaceAttributes(raceId int) (raceAttr types.RaceAttr, err error) {
	err = db.Ping()
	if err != nil {
		return raceAttr, fmt.Errorf("raceAttributePing: %v", err)
	}

	query := fmt.Sprintf(`
		SELECT  ra.*
		FROM RaceAttributes ra
		WHERE ra.RaceId = '%d'`, raceId)

	err = db.QueryRow(query).Scan(&raceAttr.Id, &raceAttr.RaceId, &raceAttr.IQ, &raceAttr.IQBonus,
		&raceAttr.ME, &raceAttr.MEBonus, &raceAttr.MA, &raceAttr.MABonus, &raceAttr.PS,
		&raceAttr.PSBonus, &raceAttr.PP, &raceAttr.PPBonus, &raceAttr.PE, &raceAttr.PEBonus,
		&raceAttr.PB, &raceAttr.PBBonus, &raceAttr.Spd, &raceAttr.SpdBonus, &raceAttr.PPE,
		&raceAttr.PPEBonus, &raceAttr.HF, &raceAttr.Alignment, &raceAttr.SpdDig, &raceAttr.SpdDigBonus)
	if err != nil {
		return raceAttr, fmt.Errorf("raceAttributeSelect: %v", err)
	}

	return raceAttr, nil
}

// Character table functions
func SaveCharacter(newChar types.Character) (id int64, err error) {
	query := fmt.Sprintf(`
		INSERT INTO palladium.Character
		(Name, RaceId, Lvl, IQ, ME, MA, PS, PP, PE, PB, Spd, HP, PPE, HF, SpdDig)
		VALUES('%s', %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d);`,
		newChar.Name, newChar.RaceId, newChar.Lvl, newChar.IQ, newChar.ME, newChar.MA, newChar.PS,
		newChar.PP, newChar.PE, newChar.PB, newChar.Spd, newChar.HP, newChar.PPE, newChar.HF, newChar.SpdDig)

	result, err := db.Exec(query)
	if err != nil {
		return 0, fmt.Errorf("newCharInsert: %v", err)
	}
	id, err = result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("newCharLastInsertId: %v", err)
	}
	return id, nil
}

func GetCharacterCount() (charCount int, err error) {
	charCount = 0
	err = db.Ping()
	if err != nil {
		return charCount, fmt.Errorf("getCharCountPing: %v", err)
	}
	query := fmt.Sprintf(`
		SELECT COUNT(c.Id)
		FROM palladium.Character c;`)

	err = db.QueryRow(query).Scan(&charCount)
	if err != nil {
		return charCount, fmt.Errorf("getCharCountSelect: %v", err)
	}
	return charCount, nil
}

func GetCharactersShort() ([]types.CharacterShort, error) {
	var characters []types.CharacterShort
	query := fmt.Sprintf(`
		SELECT c.ID, c.Name, r.Desc as Race
		FROM palladium.Character c
			JOIN palladium.Race r on r.Id = c.RaceId;`)

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("getCharnamesSelect: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var character types.CharacterShort
		if err := rows.Scan(&character.Id, &character.Name, &character.Race); err != nil {
			return nil, fmt.Errorf("getCharNamesScan: %v", err)
		}
		characters = append(characters, character)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getCharnamesFinal: %v", err)
	}
	return characters, nil
}

func GetCharacterById(id int) (character types.Character, err error) {
	err = db.Ping()
	if err != nil {
		return character, err
	}

	query := fmt.Sprintf(`
		SELECT c.Id, c.Name, c.RaceId, r.Desc as Race, c.Lvl, c.IQ, c.ME, c.MA, c.PS, c.PP, c.PE, c.PB,
			   c.Spd, c.HP, c.PPE, c.HF, c.SpdDig, c.OccId, o.Desc as OccDesc
		FROM palladium.Character c
			JOIN palladium.Race r on r.Id = c.RaceId
			JOIN palladium.OCC o on o.Id = c.OccId
		WHERE c.Id = %d`, id)

	err = db.QueryRow(query).Scan(&character.Id, &character.Name, &character.RaceId, &character.Race, &character.Lvl,
		&character.IQ, &character.ME, &character.MA, &character.PS, &character.PP, &character.PE, &character.PB,
		&character.Spd, &character.HP, &character.PPE, &character.HF, &character.SpdDig, &character.OccId, &character.OccDesc)
	if err != nil {
		return character, fmt.Errorf("characterByIdScan: %v", err)
	}

	return character, nil
}

func GetCharacterByName(name string) (character types.Character, err error) {
	err = db.Ping()
	if err != nil {
		return character, err
	}

	query := fmt.Sprintf(`
		SELECT c.Id, c.Name, c.RaceId, r.Desc as Race, c.Lvl, c.IQ, c.ME, c.MA, c.PS, c.PP, c.PE, c.PB,
			c.Spd, c.HP, c.PPE, c.HF, c.SpdDig, c.OccId, o.Desc as OccDesc
		FROM palladium.Character c
			JOIN palladium.Race r on r.Id = c.RaceId
			JOIN palladium.OCC o on o.Id = c.OccId
		WHERE c.Name = '%s'`, name)

	err = db.QueryRow(query).Scan(&character.Id, &character.Name, &character.RaceId, &character.Race, &character.Lvl,
		&character.IQ, &character.ME, &character.MA, &character.PS, &character.PP, &character.PE, &character.PB,
		&character.Spd, &character.HF, &character.PPE, &character.HF, &character.SpdDig, &character.OccId, &character.OccDesc)
	if err != nil {
		return character, fmt.Errorf("characterByIdScan: %v", err)
	}

	return character, nil
}

func GetOccsByRace(raceId int) (occs []types.Occ, err error) {

	query := fmt.Sprintf(`
		SELECT o.Id, ot.Desc as Type, o.Desc
		FROM palladium.Race_OCC ro
			LEFT JOIN palladium.OCC o on o.Id = ro.OccId 
			JOIN palladium.OCCType ot on ot.Id = o.OCCTypeId 
		WHERE ro.RaceId = %d
		ORDER BY Type, o.Desc;`, raceId)

	rows, err := db.Query(query)
	if err != nil {
		return occs, fmt.Errorf("getOccsByRaceSelect: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var occ types.Occ
		if err := rows.Scan(&occ.Id, &occ.Type, &occ.Desc); err != nil {
			return occs, fmt.Errorf("getOccsByRaceScan: %v", err)
		}
		occs = append(occs, occ)
	}
	if err := rows.Err(); err != nil {
		return occs, fmt.Errorf("getOccsByRaceFinal: %v", err)
	}
	return occs, nil
}

func SaveOcc(characterId int64, occId int) (int64, error) {
	query := fmt.Sprintf(`
		UPDATE palladium.Character
		SET OccId=%d
		WHERE Id=%d;`, occId, characterId)

	_, err := db.Exec(query)
	if err != nil {
		fmt.Printf("updateChar: %v", err)
		return 0, fmt.Errorf("updateChar: %v", err)
	}
	return characterId, nil
}
