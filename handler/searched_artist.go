package handler

import (
	"groupie_upgrade/data"
	error "groupie_upgrade/error"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func Searched_artist(w http.ResponseWriter, r *http.Request, input string) {

	var searched_data []data.All_Data
	var clear_searched_data []data.All_Data

	for _, d := range data.All_Data_Export {
		if strings.Contains(strings.ToLower(d.Name), strings.ToLower(input)) ||
			d.Firstalbum == input ||
			strconv.Itoa(d.Creationdate) == input {
			searched_data = append(searched_data, d)
		} else {
			for _, member := range d.Members {
				if strings.Contains(strings.ToLower(member), strings.ToLower(input)) {
					searched_data = append(searched_data, d)
					break
				}
			}
			for _, location := range d.Locations {
				if strings.Contains(strings.ToLower(location), strings.ToLower(input)) {
					searched_data = append(searched_data, d)
					break
				}
			}
			for _, date := range d.Dates {
				if strings.Contains(strings.ToLower(date), strings.ToLower(input)) {
					searched_data = append(searched_data, d)
					break
				}
			}
		}
	}

	if len(searched_data) > 0 {
		clear_searched_data = append(clear_searched_data, searched_data[0])
		for _, k := range searched_data {
			for j, l := range clear_searched_data {
				if k.ID == l.ID {
					break
				} else if j == len(clear_searched_data)-1 {
					clear_searched_data = append(clear_searched_data, k)
				}
			}
		}
	}

	tmpl, err := template.ParseFiles("templates/search_bar_res.html")
	if err != nil {
		error.Error_handler(w, r, 500)
		return
	}
	err = tmpl.Execute(w, clear_searched_data)
	if err != nil {
		error.Error_handler(w, r, 500)
	}
	searched_data = nil
}
