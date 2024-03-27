package handler

import (
	data "groupie_upgrade/data"
	"html/template"
	"net/http"
	"sort"
)

var Filtered_locations []string

func Home(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//****** Gestion proposition Locations filter ******

	var allLocations []string
	for _, data := range data.All_Data_Export {
		allLocations = append(allLocations, data.Locations...)
	}

	if len(allLocations) > 0 {
		Filtered_locations = append(Filtered_locations, allLocations[0])
		for _, k := range allLocations {
			for j, l := range Filtered_locations {
				if k == l {
					break
				} else if j == len(Filtered_locations)-1 {
					Filtered_locations = append(Filtered_locations, k)
				}
			}
		}
	}
	sort.Strings(Filtered_locations)

	// Exec avec All_Data + sorted locations sans doublon
	err = tmpl.Execute(w, struct {
		Data         []data.All_Data
		Filtered_loc []string
	}{
		Data:         data.All_Data_Export,
		Filtered_loc: Filtered_locations,
	})
	if err != nil {
		return
	}
}
