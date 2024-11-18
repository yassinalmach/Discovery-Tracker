package handlers

import (
	"groupie/utils"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.PageNotFound(w, r)
		return
	}

	// parseTemplate(w, "home.html", data)
}
