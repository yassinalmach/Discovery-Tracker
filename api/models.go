package api

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
