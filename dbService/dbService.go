package dbService

import (
	"database/sql"
	"fmt"
	"log"

	"pcg/attributes"

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
	fmt.Println("Connected!")
}


func GetRaces() ([]attributes.Race, error) {
	var races []attributes.Race

	rows, err := DB.Query("SELECT * FROM Race")
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