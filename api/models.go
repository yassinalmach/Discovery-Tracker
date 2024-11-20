package api

import "sync"

const URL = "https://groupietrackers.herokuapp.com/api/"

type Artists struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Location     Locations
	Date         Dates
	Relation     Relations
}

// location ***********************************
type Locations struct {
	Location []string `json:"locations"`
}
type LocationsData struct {
	Index []Locations
}

// dates ***********************************
type Dates struct {
	Date []string `json:"dates"`
}
type DatesData struct {
	Index []Dates
}

// relation ***********************************
type Relations struct {
	Relation map[string][]string `json:"datesLocations"`
}
type RelationsData struct {
	Index []Relations
}

var (
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
		fetchData(URL+"artists", &AllArtists)
	}()

	go func() {
		defer wg.Done()
		fetchData(URL+"locations", &AllLocations)
	}()

	go func() {
		defer wg.Done()
		fetchData(URL+"dates", &AllDates)
	}()

	go func() {
		defer wg.Done()
		fetchData(URL+"relation", &AllRelations)
	}()
	wg.Wait()

	for i := 0; i < 52; i++ {
		AllArtists[i].Location = AllLocations.Index[i]
		AllArtists[i].Date = AllDates.Index[i]
		AllArtists[i].Relation = AllRelations.Index[i]
	}
}
