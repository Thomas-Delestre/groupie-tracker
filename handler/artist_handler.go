package handler

import (
	data "groupie_upgrade/data"
	error "groupie_upgrade/error"
	"html/template"
	"net/http"
	"strconv"
)

func Artist_handler(w http.ResponseWriter, r *http.Request) {

	var selected_Artist_Export data.All_Data

	// recup artist ID
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		error.Error_handler(w, r, 400)
		return
	}
	if id > 52 {
		error.Error_handler(w, r, 400)
		return
	} else {
		for _, k := range data.All_Data_Export {
			if k.ID == id {
				selected_Artist_Export = k //recup la struct à l'ID identique dans S_A_E
			}
		}
	}
	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		error.Error_handler(w, r, 500)
		return
	}
	err = tmpl.Execute(w, selected_Artist_Export)
	if err != nil {
		error.Error_handler(w, r, 500)
		return
	}
}
