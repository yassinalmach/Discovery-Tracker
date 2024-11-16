package handlers

import (
	"groupie/utils"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.PageNotFound(w,r, http.StatusNotFound)
		return
	}
}