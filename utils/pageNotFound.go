package utils

import (
	"html/template"
	"net/http"
)

func PageNotFound(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./templates/notFoundPage.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	if err := tmp.Execute(w, "Error 404 (Not Found)!!"); err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}
