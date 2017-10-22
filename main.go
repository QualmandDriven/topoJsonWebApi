package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// PrepareSqLite3()
	// ImportDataFromCsv()

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func ImportDataFromCsv() {
	path := "C:\\Users\\Alexander\\Downloads\\Downloads\\GeoData"

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	database := GetSqLite3Connection()

	defer database.Close()

	for _, fInfo := range files {
		tx, _ := database.Begin()
		fmt.Println("Importing ", path+"\\"+fInfo.Name())

		f, _ := os.Open(path + "\\" + fInfo.Name())

		r := csv.NewReader(bufio.NewReader(f))
		r.Comma = '\t'
		r.LazyQuotes = true

		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				// log.Fatal(err)
				fmt.Println(fInfo.Name(), err)
				continue
			}

			statement, _ := tx.Prepare("INSERT INTO cities (id, name, asciiname, alternatenames, latitude, longitude, featureclass, featurecode, countrycode, cc2, admin1code, admin2code, admin3code, admin4code, population, elevation, dem,timezone, modificationdate) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
			statement.Exec(record[0], record[1], record[2], record[3], record[4], record[5], record[6], record[7], record[8], record[9], record[10], record[11], record[12], record[13], record[14], record[15], record[16], record[17], record[18])

			// for value := range record {
			// 	fmt.Printf("%v = %v\n", value, record[value])
			// }
			// fmt.Println(fInfo.Name(), record)
		}

		tx.Commit()
	}

}
