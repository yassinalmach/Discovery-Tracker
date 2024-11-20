package handlers

import (
	"groupie/api"
	"groupie/utils"
	"net/http"
	"strconv"
)

type Data struct {
	Artist []api.Artists
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.PageNotFound(w, r)
		return
	}
	// init data
	// data := Data{
	// 	Artist: api.AllArtists,
	// }

	// filter section
	var filteredData []api.Artists

	if err := filterData(r, &filteredData); err != nil {
		http.Error(w, "bad request", 400)
	}

	parseTemplate(w, "home.html", filteredData)
}

func filterData(r *http.Request, filteredData *[]api.Artists) error {
	r.ParseForm()
	nbrMembers := r.Form["nbrMembers"]
	minValueDate := r.FormValue("minValueDate")
	maxValueDate := r.FormValue("maxValueDate")
	minValueAlbum := r.FormValue("minValueAlbum")
	maxValueAlbum := r.FormValue("maxValueAlbum")
	locationValue := r.FormValue("location")

	if len(nbrMembers) == 0 && minValueDate == "" && maxValueDate == "" &&
		minValueAlbum == "" && maxValueAlbum == "" && locationValue == "" {
		filteredData = &api.AllArtists
		return nil
	}

	for _, artist := range api.AllArtists {
		var isValid bool
		var err error
		// filter by members
		isValid, err = filterMembers(filteredData, artist, nbrMembers)
		if !isValid {
			continue
		}

		err, isValid = filterCreationDate(filteredData, artist, minValueDate, maxValueDate)
		if !isValid {
			continue
		}

		err, isValid = filterFirstAlbum(filteredData, artist, minValueAlbum, maxValueAlbum)
		if !isValid {
			continue
		}

		err, isValid = filterLocation(filteredData, artist, locationValue)
		if !isValid {
			continue
		}

		if err != nil {
			return err
		}

		if isValid {
			*filteredData = append(*filteredData, artist)
		}

	}
	return nil
}

func filterMembers(filteredData *[]api.Artists, artist api.Artists, nbrMembers []string) (bool, error) {
	if len(nbrMembers) != 0 {
		for _, v := range nbrMembers {
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
	return true, nil
}

func filterCreationDate(filteredData *[]api.Artists, artist api.Artists, minValueDate, maxValueDate string) (error, bool) {
	return nil, true
}

func filterFirstAlbum(filteredData *[]api.Artists, artist api.Artists, minValueAlbum, maxValueAlbum string) (error, bool) {
	return nil, true
}

func filterLocation(filteredData *[]api.Artists, artist api.Artists, locationValue string) (error, bool) {
	return nil, true
}
