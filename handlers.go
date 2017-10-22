package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintf(w, "Todo show: %s\n", todoId)
}

func GetCountries(w http.ResponseWriter, r *http.Request) {
	countries := Countries{
		Country{Name: "Germany"},
		Country{Name: "Germany"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(countries)
}

func GetCities(w http.ResponseWriter, r *http.Request) {

	database := GetSqLite3Connection()
	defer database.Close()

	rows, _ := database.Query("SELECT id, name FROM cities WHERE name='Munich'")

	cities := Cities{}

	var id int
	var name string
	for rows.Next() {
		rows.Scan(&id, &name)
		cities = append(cities, City{Id: id, Name: name})
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(cities); err != nil {
		panic(err)
	}
}
