package main

import (
	"database/sql"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

const concurrency = 5

func main() {
	db, err := sql.Open("sqlite", "./cities.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT latitude, longitude FROM cities WHERE city = ? ORDER BY population DESC LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}

	// some lines are commented out because they are cities that don't exist
	// uncomment them to witness the mmap leakage
	queries := []struct {
		city string
	}{
		{"berlin"},
		{"san francisco"},
		{"london"},
		{"chicago"},
		{"houston"},
		{"los angeles"},
		{"new york"},
		{"london"},
		{"chicago"},
		{"houston"},
		{"dublin"},
		{"seattle"},
		{"portland"},
		{"alabama"},
		{"dallas"},
		{"brooklyn"},
		{"philadelphia"},
		{"atlanta"},
		{"phoenix"},
		{"hanoi"},
		{"abu dhabi"},
		{"abuja"},
		{"accra"},
		{"adamstown"},
		{"addis ababa"},
		{"algiers"},
		{"alofi"},
		{"amman"},
		{"amsterdam"},
		{"andorra la vella"},
		{"ankara"},
		{"antananarivo"},
		{"apia"},
		{"ashgabat"},
		{"asmara"},
		{"athens"},
		{"avarua"},
		{"baghdad"},
		{"baku"},
		{"bamako"},
		{"bandar seri begawan"},
		{"bangkok"},
		{"bangui"},
		{"banjul"},
		{"basseterre"},
		{"beijing"},
		{"beirut"},
		{"belgrade"},
		{"belmopan"},
		{"berlin"},
		{"bern"},
		{"bishkek"},
		{"bissau"},
		{"bratislava"},
		{"brazzaville"},
		{"bridgetown"},
		{"brussels"},
		{"bucharest"},
		{"budapest"},
		{"buenos aires"},
		{"cairo"},
		{"canberra"},
		{"caracas"},
		{"castries"},
		{"charlotte amalie"},
		{"does not exist"},
		{"cockburn town"},
		{"conakry"},
		{"copenhagen"},
		{"dakar"},
		{"damascus"},
		{"dhaka"},
		{"dili"},
		{"djibouti"},
		{"dodoma"},
		{"doha"},
		{"douglas"},
		{"dublin"},
		{"dushanbe"},
	}

	for {
		for g := 0; g < concurrency; g += 1 {
			go func() {
				for _, query := range queries {
					var long, lat float32
					row := stmt.QueryRow(query.city)

					if err := row.Scan(&lat, &long); err != nil {
						log.Printf("error for city %q: %#v\n", query.city, err)
					}
				}
			}()
		}

		log.Println("batch done")

		// Give a small break between iterations to avoid consuming all system memory too quickly
		time.Sleep(3 * time.Second)
	}
}
