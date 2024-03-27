package error

import (
	"fmt"
	"net/http"
	"text/template"
)

func Error_handler(w http.ResponseWriter, r *http.Request, Err_status int) {

	var verif bool
	var tmpl *template.Template
	// on remarque que template.Parsefiles a besoin d'une structure pointé
	// on l'initie avant car on va l'executer en dehors du switch/case
	// w.WriteHeader peut être mis au début comme à la fin
	var err error

	switch Err_status {
	case 400: // StatusBadRequest
		w.WriteHeader(400)
		tmpl, err = template.ParseFiles("templates/error_400.html")
		// on peut aussi simplement le lire avec http.ServeFile(w, r, "page_html/error_404.html")
		fmt.Println("Error 400")
		verif = true

	case 404: // StatusNotFound
		w.WriteHeader(404)
		tmpl, err = template.ParseFiles("templates/error_404.html")
		fmt.Println("Error 404")
		verif = true

	case 500: // StatusInternalServerError
		w.WriteHeader(500)
		tmpl, err = template.ParseFiles("templates/error_500.html")
		fmt.Println("Error 500")
		verif = true

	default:
		tmpl, err = template.ParseFiles("index.html")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !verif {
		w.WriteHeader(Err_status)
		// Permet d'afficher le status d'une erreur que l'on ne gère pas avec cette fonction
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), 400)
		// Error_Handler(w, r, 400) => on ne peut pas faire
		return
	}
}
