package api

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

const URL = "https://groupietrackers.herokuapp.com/api/"

type Artists struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

// location ***********************************
type Locations struct {
	Location []string `json:"locations"`
}

type LocationsData struct {
	Index []struct {
		Location []string `json:"locations"`
	}
}

// dates ***********************************
type Dates struct {
	Date []string `json:"dates"`
}
type DatesData struct {
	Index []struct {
		Date []string `json:"dates"`
	}
}

// relation ***********************************
type Relations struct {
	Relation map[string][]string `json:"datesLocations"`
}
type RelationsData struct {
	Index []struct {
		Relation map[string][]string `json:"datesLocations"`
	}
}

type Data struct {
	Artist   Artists
	Location Locations
	Date     Dates
	Relation Relations
}

var (
	AllData      []Data
	AllArtists   []Artists
	AllLocations LocationsData
	AllDates     DatesData
	AllRelations RelationsData
)

func init() {
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		FetchData(URL+"artists", &AllArtists)
	}()

	go func() {
		defer wg.Done()
		FetchData(URL+"locations", &AllLocations)
	}()

	go func() {
		defer wg.Done()
		FetchData(URL+"dates", &AllDates)
	}()

	go func() {
		defer wg.Done()
		FetchData(URL+"relations", &AllRelations)
	}()
	wg.Wait()

	for i := 0; i < 52; i++ {
		data := Data{
			Artist:   AllArtists[i],
			Location: AllLocations.Index[i],
			Date:     AllDates.Index[i],
			Relation: AllRelations.Index[i],
		}
		AllData = append(AllData, data)
	}
}

func FetchData(url string, data any) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(data)
}
