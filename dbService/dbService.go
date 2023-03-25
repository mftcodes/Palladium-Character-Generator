package dbService

import (
	"database/sql"
	"fmt"
	"log"

	"pcg/attributes"
	"pcg/character"

    "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DbConnect() {
	fmt.Println("Connecting to Db")
	cfg := mysql.Config{
		User:   "root",
		Passwd: "pcg0superfun",
		Net:    "tcp",
		Addr:   "localhost:3308",
		DBName: "palladium",
	}
	// Get a database handle.
	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected and ready to begin!\n")
}


func GetRaces() ([]attributes.Race, error) {
	var races []attributes.Race
	query := `SELECT * FROM Race`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("racesSelect: %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var race attributes.Race
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

func GetRaceAttributes(raceId int) (attributes.RaceAttributes, error) {
	err := DB.Ping()
	if err != nil {
		return attributes.RaceAttributes{}, err
	}
	
	var raceAttributes attributes.RaceAttributes
	query := fmt.Sprintf(`
		SELECT  ra.*
		FROM RaceAttributes ra
		WHERE ra.RaceId =  + %d`, raceId)

	err = DB.QueryRow(query).Scan(&raceAttributes.Id, &raceAttributes.RaceId, &raceAttributes.IQ, &raceAttributes.IQBonus, 
		&raceAttributes.ME, &raceAttributes.MEBonus, &raceAttributes.MA,  &raceAttributes.MABonus, &raceAttributes.PS, 
		&raceAttributes.PSBonus, &raceAttributes.PP, &raceAttributes.PPBonus, &raceAttributes.PE, &raceAttributes.PEBonus, 
		&raceAttributes.PB, &raceAttributes.PBBonus, &raceAttributes.Spd, &raceAttributes.SpdBonus, &raceAttributes.PPE, 
		&raceAttributes.PPEBonus, &raceAttributes.Alignment, &raceAttributes.SpdDig, &raceAttributes.SpdDigBonus)
	if err != nil {
		return raceAttributes, fmt.Errorf("raceAttributeSelect: %v", err)
	}

	// raceAttributes = row
	return raceAttributes, nil
}

func SaveCharacter(newChar character.Character) (int64, error) {
	query := fmt.Sprintf(`
		INSERT INTO palladium.Character
		(Name, RaceId, Lvl, IQ, ME, MA, PS, PP, PE, PB, Spd, PPE, SpdDig)
		VALUES('%s', %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d);`,
		newChar.Name, newChar.RaceId, newChar.Lvl, newChar.IQ, newChar.ME, newChar.MA, newChar.PS, 
		newChar.PP, newChar.PE, newChar.PB, newChar.Spd, newChar.PPE, newChar.SpdDig)

	fmt.Println(query)
	result, err :=  DB.Exec(query)
	if err != nil {
		return 0, fmt.Errorf("newCharInsert: %v", err)
	}
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("newCharLastInsertId: %v", err)
    }
    return id, nil
}