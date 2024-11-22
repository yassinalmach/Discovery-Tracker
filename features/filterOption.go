package features

import (
	"groupie/api"
	"strconv"
	"strings"
)

type Filter struct {
	NbrMembers    []string
	MinValueDate  string
	MaxValueDate  string
	MinValueAlbum string
	MaxValueAlbum string
	LocationValue string
}

// filterData filters artists
func FilterData(f Filter) ([]api.Artists, error) {
	// If no filters applied, return all data
	if len(f.NbrMembers) == 0 && f.MinValueDate == "" && f.MaxValueDate == "" &&
		f.MinValueAlbum == "" && f.MaxValueAlbum == "" && f.LocationValue == "" {
		return api.AllArtists, nil
	}

	var filteredData []api.Artists

	// Iterate over artists and apply filters
	for _, artist := range api.AllArtists {
		var isValid bool
		var err error

		// filter by members
		if isValid, err = filterMembers(artist, f.NbrMembers); err != nil {
			return nil, err
		} else if !isValid {
			continue
		}

		// filter by creation date
		if isValid, err = filterByDate(artist.CreationDate, f.MinValueDate, f.MaxValueDate); err != nil {
			return nil, err
		} else if !isValid {
			continue
		}

		// filter by first album
		albumDate, err := strconv.Atoi(strings.Split(artist.FirstAlbum, "-")[2])
		if err != nil {
			return nil, err
		}

		if isValid, err = filterByDate(albumDate, f.MinValueAlbum, f.MaxValueAlbum); err != nil {
			return nil, err
		} else if !isValid {
			continue
		}

		// filter by location
		if isValid = filterLocation(artist, f.LocationValue); !isValid {
			continue
		}

		if isValid {
			filteredData = append(filteredData, artist)
		}
	}
	return filteredData, nil
}

// filterMembers checks if artist has matching number of members
func filterMembers(artist api.Artists, NbrMembers []string) (bool, error) {
	if len(NbrMembers) == 0 {
		return true, nil
	}

	for _, v := range NbrMembers {
		nbr, err := strconv.Atoi(v)
		if err != nil {
			return false, err
		}
		if len(artist.Members) == nbr {
			return true, nil
		}
	}
	return false, nil
}

// filterByDate checks if date is within range
func filterByDate(year int, min, max string) (bool, error) {
	var err error
	if min == "" && max == "" {
		return true, nil
	}

	minYear := 1950
	maxYear := 2024

	if min != "" {
		minYear, err = strconv.Atoi(min)
		if err != nil {
			return false, err
		}
	}

	if max != "" {
		maxYear, err = strconv.Atoi(max)
		if err != nil {
			return false, err
		}
	}

	return year >= minYear && year <= maxYear, nil
}

// filterLocation checks if location contains the search value
func filterLocation(artist api.Artists, value string) bool {
	if value == "" {
		return true
	}

	for _, loc := range artist.Location.Location {
		if strings.Contains(loc, value) {
			return true
		}
	}
	return false
}
