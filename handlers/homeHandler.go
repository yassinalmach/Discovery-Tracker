package handlers

import (
	"groupie/api"
	"groupie/utils"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.PageNotFound(w, r)
		return
	}

	var data []api.Artists
	api.FetchData(w, api.URL+"/artists", &data)
	parseTemplate(w, "index.html", data)
}
