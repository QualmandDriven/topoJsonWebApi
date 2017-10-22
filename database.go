package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetSqLite3Connection() *sql.DB {
	database, _ := sql.Open("sqlite3", "./geo.db")

	return database
}

func PrepareSqLite3() {
	database := GetSqLite3Connection()
	defer database.Close()
	statement, _ := database.Prepare(`CREATE TABLE IF NOT EXISTS cities (id INTEGER PRIMARY KEY,
																		 name TEXT,
																		 asciiname TEXT,
																		 alternatenames TEXT,
																		 latitude REAL,
																		 longitude REAL,
																		 featureclass TEXT,
																		 featurecode TEXT,
																		 countrycode TEXT,
																		 cc2 TEXT,
																		 admin1code TEXT,
																		 admin2code TEXT,
																		 admin3code TEXT,
																		 admin4code TEXT,
																		 population INTEGER,
																		 elevation INTEGER,
																		 dem TEXT,
																		 timezone TEXT,
																		 modificationdate TEXT)`)
	statement.Exec()
	// statement, _ = database.Prepare("INSERT INTO people (id, name) VALUES (?, ?)")
	// statement.Exec("Nic", "Raboy")
	// rows, _ := database.Query("SELECT id, firstname, lastname FROM people")
	// var id int
	// var firstname string
	// var lastname string
	// for rows.Next() {
	// 	rows.Scan(&id, &firstname, &lastname)
	// 	fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	// }
}
