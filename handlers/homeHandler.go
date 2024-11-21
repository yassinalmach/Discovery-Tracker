package handlers

import (
	"groupie/features"
	"groupie/utils"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.PageNotFound(w, r)
		return
	}

	// filter section
	var filter features.Filter

	// parse values
	r.ParseForm()
	filter.NbrMembers = r.Form["nbrMembers"]
	filter.MinValueDate = r.FormValue("minValueDate")
	filter.MaxValueDate = r.FormValue("maxValueDate")
	filter.MinValueAlbum = r.FormValue("minValueAlbum")
	filter.MaxValueAlbum = r.FormValue("maxValueAlbum")
	filter.LocationValue = r.FormValue("location")

	filteredData, err := features.FilterData(filter)
	if err != nil {
		http.Error(w, "bad request", 400)
	}

	parseTemplate(w, "home.html", filteredData)
}
