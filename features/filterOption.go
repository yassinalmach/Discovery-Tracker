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
	// check if no filter option, it returns all data
	if len(f.NbrMembers) == 0 && f.MinValueDate == "" && f.MaxValueDate == "" &&
		f.MinValueAlbum == "" && f.MaxValueAlbum == "" && f.LocationValue == "" {
		return api.AllArtists, nil
	}

	var filteredData []api.Artists

	// iterate over artists
	for _, artist := range api.AllArtists {
		var isValid bool
		var err error
		// filter by members
		isValid, err = filterMembers(artist, f.NbrMembers)
		if err != nil {
			return nil, err
		} else if !isValid {
			continue
		}

		// filter by creation date
		isValid, err = filterCreationDate(artist, f.MinValueDate, f.MaxValueDate)
		if err != nil {
			return nil, err
		} else if !isValid {
			continue
		}

		// filter by first album
		isValid, err = filterFirstAlbum(artist, f.MinValueAlbum, f.MaxValueAlbum)
		if err != nil {
			return nil, err
		} else if !isValid {
			continue
		}

		// filter by location
		isValid, err = filterLocation(artist, f.LocationValue)
		if err != nil {
			return nil, err
		} else if !isValid {
			continue
		}

		if isValid {
			filteredData = append(filteredData, artist)
		}
	}
	return filteredData, nil
}

// filterMembers filter artists by the number of members
func filterMembers(artist api.Artists, NbrMembers []string) (bool, error) {
	if len(NbrMembers) == 0 {
		return true, nil
	} else {
		for _, v := range NbrMembers {
			nbr, err := strconv.Atoi(v)
			if err != nil {
				return false, err
			}
			if len(artist.Members) == nbr {
				return true, nil
			}
		}
	}
	return false, nil
}

// filterCreationDate filter artists by creation date
func filterCreationDate(artist api.Artists, min, max string) (bool, error) {
	if min == "" && max == "" {
		return true, nil
	}

	if min != "" {
		// convert min to number
		minNbr, err := strconv.Atoi(min)
		if err != nil {
			return false, err
		}

		// if min and max exists
		if max != "" {
			// convert max to a number
			maxNbr, err := strconv.Atoi(max)
			if err != nil {
				return false, err
			}

			if minNbr > maxNbr {
				return false, err
			}

			if artist.CreationDate >= minNbr && artist.CreationDate <= maxNbr {
				return true, nil
			}

			// if only have a min number
		} else {
			if artist.CreationDate >= minNbr {
				return true, nil
			}
		}

		// if only have max number
	} else if max != "" {
		maxNbr, err := strconv.Atoi(max)
		if err != nil {
			return false, err
		}
		if artist.CreationDate <= maxNbr {
			return true, nil
		}
	}
	return false, nil
}

// filterFirstAlbum filter artists by first album's date
func filterFirstAlbum(artist api.Artists, min, max string) (bool, error) {
	firstAlbum := strings.Split(artist.FirstAlbum, "-")[2]

	if min == "" && max == "" {
		return true, nil
	}

	if min != "" {
		if max != "" {
			if firstAlbum >= min && firstAlbum <= max {
				return true, nil
			}
		} else {
			if firstAlbum >= min {
				return true, nil
			}
		}
	} else {
		if firstAlbum <= max {
			return true, nil
		}
	}
	return false, nil
}

// filterLocation filter artists by concert location
func filterLocation(artist api.Artists, value string) (bool, error) {
	if value == "" {
		return true, nil
	}

	for _, loc := range artist.Location.Location {
		if strings.Contains(loc, value) {
			return true, nil
		}
	}
	return false, nil
}
