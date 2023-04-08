package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Connect
func dbConnect() {
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
func getRaces() ([]race, error) {
	var races []race
	query := `SELECT * FROM Race`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("racesSelect: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var race race
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

func getRaceByName(name string) (race, error) {
	err := db.Ping()
	if err != nil {
		return race{}, err
	}

	var race race
	query := fmt.Sprintf(`SELECT * FROM Race r WHERE r.Desc = %s`, name)

	err = db.QueryRow(query).Scan(&race.Id, &race.Desc)
	if err != nil {
		return race, fmt.Errorf("raceSelectByName: %v", err)
	}

	return race, nil
}

func getRaceById(id int) (race, error) {
	err := db.Ping()
	if err != nil {
		return race{}, err
	}

	var race race
	query := fmt.Sprintf(`SELECT * FROM Race r WHERE r.Id = %d`, id)

	err = db.QueryRow(query).Scan(&race.Id, &race.Desc)
	if err != nil {
		return race, fmt.Errorf("raceSelectByName: %v", err)
	}

	return race, nil
}

func getRaceAttributes(raceId int) (raceAttributes, error) {
	err := db.Ping()
	if err != nil {
		return raceAttributes{}, fmt.Errorf("raceAttributePing: %v", err)
	}

	var raceAttributes raceAttributes
	query := fmt.Sprintf(`
		SELECT  ra.*
		FROM RaceAttributes ra
		WHERE ra.RaceId =  + %d`, raceId)

	err = db.QueryRow(query).Scan(&raceAttributes.Id, &raceAttributes.RaceId, &raceAttributes.IQ, &raceAttributes.IQBonus,
		&raceAttributes.ME, &raceAttributes.MEBonus, &raceAttributes.MA, &raceAttributes.MABonus, &raceAttributes.PS,
		&raceAttributes.PSBonus, &raceAttributes.PP, &raceAttributes.PPBonus, &raceAttributes.PE, &raceAttributes.PEBonus,
		&raceAttributes.PB, &raceAttributes.PBBonus, &raceAttributes.Spd, &raceAttributes.SpdBonus, &raceAttributes.PPE,
		&raceAttributes.PPEBonus, &raceAttributes.HF, &raceAttributes.Alignment, &raceAttributes.SpdDig, &raceAttributes.SpdDigBonus)
	if err != nil {
		return raceAttributes, fmt.Errorf("raceAttributeSelect: %v", err)
	}

	return raceAttributes, nil
}

// Character table functions
func saveCharacter(newChar character) (int64, error) {
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
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("newCharLastInsertId: %v", err)
	}
	return id, nil
}

func getCharacterCount() (int, error) {
	charCount := 0
	err := db.Ping()
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

func getCharactersShort() ([]characterShort, error) {
	var characters []characterShort
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
		var character characterShort
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

func getCharacterById(id int) (character, error) {
	var character character
	err := db.Ping()
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

func getCharacterByName(name string) (character, error) {
	var character character
	err := db.Ping()
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

func getOccsByRace(raceId int) ([]occ, error) {
	var occs []occ

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
		var occ occ
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

func saveOcc(characterId int64, occId int) (int64, error) {
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
