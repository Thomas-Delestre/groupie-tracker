package data

import (
	"encoding/json"
	"log"
	"net/http"
)

var All_Data_Export []All_Data

func FetchData(apiURL string) {
	response, err := http.Get(apiURL)
	if err != nil {
		log.Fatal(err)
		// Gestion de l'erreur selon vos besoins
	}

	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)

	var artist_data []Artist_data

	if err := decoder.Decode(&artist_data); err != nil {
		log.Fatal(err)
		// Gestion de l'erreur
	}

	for _, k := range artist_data {

		// récupération donnée Locations_data
		response, err := http.Get(k.Locations)
		if err != nil {
			log.Fatal(err)
			// Gestion de l'erreur selon vos besoins
		}
		defer response.Body.Close()

		decoder := json.NewDecoder(response.Body)

		var locations_data Locations_data

		if err := decoder.Decode(&locations_data); err != nil {
			log.Fatal(err)
			// Gestion de l'erreur selon vos besoins
		}

		// récupération donnée Concert_dates_data
		response, err = http.Get(k.Concert_dates)
		if err != nil {
			log.Fatal(err)
			// Gestion de l'erreur selon vos besoins
		}

		defer response.Body.Close()

		decoder = json.NewDecoder(response.Body)

		var concert_dates_data Concert_dates_data

		if err := decoder.Decode(&concert_dates_data); err != nil {
			log.Fatal(err)
			// Gestion de l'erreur selon vos besoins
		}

		// récupération donnée Relations_data
		response, err = http.Get(k.Relations)
		if err != nil {
			log.Fatal(err)
			// Gestion de l'erreur selon vos besoins
		}

		defer response.Body.Close()

		decoder = json.NewDecoder(response.Body)

		var relations_data Relations_data

		if err := decoder.Decode(&relations_data); err != nil {
			log.Fatal(err)
			// Gestion de l'erreur selon vos besoins
		}

		All_Data_Export = append(All_Data_Export, All_Data{
			ID:             k.ID,
			Name:           k.Name,
			Image:          k.Image,
			Members:        k.Members,
			Creationdate:   k.Creationdate,
			Firstalbum:     k.Firstalbum,
			Locations:      locations_data.Locations,
			Dates:          concert_dates_data.Dates,
			DatesLocations: relations_data.DatesLocations,
		})
	}
}

// decoder := json.NewDecoder(response.Body)
// if err := decoder.Decode(data); err != nil {
// 	log.Fatal(err)
// 	// Gestion de l'erreur selon vos besoins
// }
