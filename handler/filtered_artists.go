package handler

import (
	data "groupie_upgrade/data"
	error "groupie_upgrade/error"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
)

func Filtered_artists(w http.ResponseWriter, r *http.Request) {

	// recup querry
	r.ParseForm()

	// stockage valeurs querry (voir struct.go)
	filters := data.Filters_values{
		CD_start: Get_int(r.Form.Get("CD_min_value")),
		CD_end:   Get_int(r.Form.Get("CD_max_value")),
		FA_start: Get_int(r.Form.Get("FA_min_value")),
		FA_end:   Get_int(r.Form.Get("FA_max_value")),
		Checkbox: getCheckboxValues(r.Form),
		Location: r.Form.Get("location"),
	}

	//fmt.Println(filters)

	var filtered_data []data.All_Data
	var clear_filtered_data []data.All_Data

	// cascade de filtres
	for _, k := range data.All_Data_Export {
		//recup année du First_Album en int
		first_album_year, _ := strconv.Atoi(k.Firstalbum[len(k.Firstalbum)-4 : len(k.Firstalbum)])
		//check Creation Data filter
		if k.Creationdate >= filters.CD_start && k.Creationdate <= filters.CD_end {
			//check First_Album year
			if first_album_year >= filters.FA_start && first_album_year <= filters.FA_end {
				//check nombre de membre demandé
				var box_checker int
				for _, k := range filters.Checkbox {
					box_checker += k
				}
				filou := 1
				for _, box_bool := range filters.Checkbox {

					if box_bool != 0 || box_checker == 0 {

						if len(k.Members) == filou || box_checker == 0 {

							for _, city := range k.Locations {
								if city == filters.Location || filters.Location == "All" {
									filtered_data = append(filtered_data, k)
									break
								}

							}
						}
					}
					filou++ // filouterie supprimable à tester
				}
			}
		}
	}
	// suppression des doublons, certains test de l'audit j'affichait des doublons. Donc ça peut être une filouterie
	if len(filtered_data) > 0 {
		clear_filtered_data = append(clear_filtered_data, filtered_data[0])
		for _, k := range filtered_data {
			for j, l := range clear_filtered_data {
				if k.ID == l.ID {
					break
				} else if j == len(clear_filtered_data)-1 {
					clear_filtered_data = append(clear_filtered_data, k)
				}
			}
		}
	}
	// lancement page avec struct d'artists filtrés
	tmpl, err := template.ParseFiles("templates/filtered_artists.html")
	if err != nil {
		error.Error_handler(w, r, 500)
		return
	}
	err = tmpl.Execute(w, struct {
		Data         []data.All_Data
		Filtered_loc []string
	}{
		Data:         clear_filtered_data,
		Filtered_loc: Filtered_locations,
	})

	if err != nil {
		error.Error_handler(w, r, 500)
	}
}

// pour un Atoi sans erreur et une initialisation plus simple de la valeur
func Get_int(value string) int {
	if intValue, err := strconv.Atoi(value); err == nil {
		if intValue == 1957 {
			return 0
		}
		return intValue
	}
	return 0
}

// Récup des bool des checkbox dans un tableau de 7 valeurs.
func getCheckboxValues(form url.Values) [7]int {
	checkboxNames := []string{
		"box_1_mbr",
		"box_2_mbr",
		"box_3_mbr",
		"box_4_mbr",
		"box_5_mbr",
		"box_6_mbr",
		"box_7_mbr",
	}
	var checkboxValues [7]int
	for i, name := range checkboxNames {
		if form.Get(name) == "on" {
			checkboxValues[i] = 1
		}
	}
	return checkboxValues
}
