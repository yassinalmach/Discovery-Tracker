package handlers

import (
	"fmt"
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
	query := strings.ToLower(r.FormValue("search"))

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
	var suggestion []api.SuggestionModel
	isAdded := make(map[string]bool)

	for _, artist := range api.AllArtists {
		// Add artist name
		addSuggestion(&suggestion, isAdded, "artist/band", artist.Name)

		// Add first album
		addSuggestion(&suggestion, isAdded, "first album", artist.FirstAlbum)

		// Add creation date
		addSuggestion(&suggestion, isAdded, "creation date", fmt.Sprintf("%v", artist.CreationDate))

		// Add members
		for _, member := range artist.Members {
			addSuggestion(&suggestion, isAdded, "member", member)
		}

		// Add locations
		for _, loc := range artist.Location.Location {
			addSuggestion(&suggestion, isAdded, "location", loc)
		}
	}
	return suggestion
}

func addSuggestion(result *[]api.SuggestionModel, isAdded map[string]bool, typeName, value string) {
	key := typeName + ":" + value
	if !isAdded[key] {
		*result = append(*result, api.SuggestionModel{Type: typeName, Value: value})
		isAdded[key] = true
	}
}

func searchFunc(query string) []api.Artists {
	var filteredData []api.Artists
	for _, artist := range api.AllArtists {
		if strings.Contains(strings.ToLower(artist.Name), query) ||
			strings.Contains(strings.ToLower(artist.FirstAlbum), query) ||
			strings.Contains(fmt.Sprintf("%v", artist.CreationDate), query) ||
			containsAny(artist.Members, query) || containsAny(artist.Location.Location, query) {
			filteredData = append(filteredData, artist)
		}
	}
	return filteredData
}

func containsAny(value []string, query string) bool {
	for _, v := range value {
		if strings.Contains(strings.ToLower(v), query) {
			return true
		}
	}
	return false
}
