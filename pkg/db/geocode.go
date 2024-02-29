package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var geocodeDB *sql.DB

func PromptGeocodeDB() {
	driver := "sqlite3"
	dataLoc := "./db/geocode.db"
	initStatement := `
		CREATE TABLE IF NOT EXISTS
		geocode (
			area TEXT,
			latitude REAL,
			longitude REAL
		)
	`

	database, err := sql.Open(driver, dataLoc)
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.Exec(initStatement)
	if err != nil {
		log.Fatal(err)
	}

	geocodeDB = database
}

func QueryGeocode(area string) (float64, float64, bool) {
	var latitude, longitude float64

	queryStatement := "SELECT latitude, longitude FROM geocode WHERE area = ?"
	row := geocodeDB.QueryRow(queryStatement, area)
	err := row.Scan(&latitude, &longitude)
	if err == sql.ErrNoRows {
		return 0, 0, false
	}

	return latitude, longitude, true
}

func AddNewGeocode(area string, latitude, longitude float64) {
	ps, err := geocodeDB.Begin()
	if err != nil {
		log.Fatal(err)
	}

	insertStatement := "INSERT INTO geocode(area, latitude, longitude) VALUES(?,?,?)"
	statement, err := geocodeDB.Prepare(insertStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	_, err = statement.Exec(area, latitude, longitude)
	if err != nil {
		log.Fatal(err)
	}

	err = ps.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
