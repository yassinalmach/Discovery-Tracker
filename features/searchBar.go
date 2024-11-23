package features

import (
	"fmt"
	"groupie/api"
	"strings"
)

// GenerateSuggetions generates suggestions to send it to html
func GenerateSuggetions() []api.SuggestionModel {
	var suggestion []api.SuggestionModel
	isAdded := make(map[string]bool)

	for _, artist := range api.AllArtists {
		// Add artist name
		addSuggestion(&suggestion, isAdded, "artist/band", artist.Name)

		// Add first album
		addSuggestion(&suggestion, isAdded, "first-album", artist.FirstAlbum)

		// Add creation date
		addSuggestion(&suggestion, isAdded, "creation-date", fmt.Sprintf("%v", artist.CreationDate))

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

// SearchFunc matches query with artists
func SearchFunc(query string) []api.Artists {
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
