package handlers

import (
	"groupie/api"
	"groupie/features"
	"groupie/utils"
	"net/http"
	"strings"
)

type Data struct {
	Artists     []api.Artists
	Suggestions []api.SuggestionModel
	Query       string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.PageNotFound(w, r)
		return
	}

	// ****** filter section ******
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

	// ****** search section ******
	query := strings.ToLower(r.URL.Query().Get("search"))
	if query != "" {
		filteredData = searchFunc(query)
	}

	suggestions := generateSuggetions()

	pageData := Data{
		Artists:     filteredData,
		Query:       query,
		Suggestions: suggestions,
	}

	parseTemplate(w, "home.html", pageData)
}

func generateSuggetions() []api.SuggestionModel {
	return nil
}

func searchFunc(query string) []api.Artists {
	return nil
}
