package main

import (
	"database/sql"
	"fmt"
	"log"

    "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func dbConnect() {
	fmt.Println("Connecting to Db")
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
	fmt.Println("Connected and ready to begin!\n")
}


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
		if err := rows.Scan(&race.Id, &race.Name); err != nil {
			return nil, fmt.Errorf("raceScan: %v", err)
		}
		races = append(races, race)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("racesFinal: %v", err)
	}
	return races, nil
}

func GetRaceByName(name string) (race, error) {
	err := db.Ping()
	if err != nil {
		return race{}, err
	}

	var race race
	query := fmt.Sprintf(`SELECT * FROM Race r WHERE r.Name = %s`, name)

	err = db.QueryRow(query).Scan(&race.Id, &race.Name)
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

	err = db.QueryRow(query).Scan(&race.Id, &race.Name)
	if err != nil {
		return race, fmt.Errorf("raceSelectByName: %v", err)
	}
	
	return race, nil
}

func getRaceAttributes(raceId int) (raceAttributes, error) {
	err := db.Ping()
	if err != nil {
		return raceAttributes{}, err
	}
	
	var raceAttributes raceAttributes
	query := fmt.Sprintf(`
		SELECT  ra.*
		FROM RaceAttributes ra
		WHERE ra.RaceId =  + %d`, raceId)

	err = db.QueryRow(query).Scan(&raceAttributes.Id, &raceAttributes.RaceId, &raceAttributes.IQ, &raceAttributes.IQBonus, 
		&raceAttributes.ME, &raceAttributes.MEBonus, &raceAttributes.MA,  &raceAttributes.MABonus, &raceAttributes.PS, 
		&raceAttributes.PSBonus, &raceAttributes.PP, &raceAttributes.PPBonus, &raceAttributes.PE, &raceAttributes.PEBonus, 
		&raceAttributes.PB, &raceAttributes.PBBonus, &raceAttributes.Spd, &raceAttributes.SpdBonus, &raceAttributes.PPE, 
		&raceAttributes.PPEBonus, &raceAttributes.Alignment, &raceAttributes.SpdDig, &raceAttributes.SpdDigBonus)
	if err != nil {
		return raceAttributes, fmt.Errorf("raceAttributeSelect: %v", err)
	}

	return raceAttributes, nil
}

func saveCharacter(newChar character) (int64, error) {
	query := fmt.Sprintf(`
		INSERT INTO palladium.Character
		(Name, RaceId, Lvl, IQ, ME, MA, PS, PP, PE, PB, Spd, PPE, SpdDig)
		VALUES('%s', %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d);`,
		newChar.Name, newChar.RaceId, newChar.Lvl, newChar.IQ, newChar.ME, newChar.MA, newChar.PS, 
		newChar.PP, newChar.PE, newChar.PB, newChar.Spd, newChar.PPE, newChar.SpdDig)

	fmt.Println(query)
	result, err :=  db.Exec(query)
	if err != nil {
		return 0, fmt.Errorf("newCharInsert: %v", err)
	}
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("newCharLastInsertId: %v", err)
    }
    return id, nil
}