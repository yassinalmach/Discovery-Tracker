package handlers

import (
	"html/template"
	"net/http"
)

func parseTemplate(w http.ResponseWriter, path string, data any) {
	tmp, err := template.ParseFiles("./templates/" + path)
	if err != nil {
		http.Error(w, "internal server error", 500)
		return
	}
	if err := tmp.Execute(w, data); err != nil {
		http.Error(w, "internal server error", 500)
		return
	}
}
