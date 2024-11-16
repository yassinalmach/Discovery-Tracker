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

type Locations struct {
	Location []string `json:"locations"`
}

type Dates struct {
	Date []string `json:"dates"`
}
type Relations struct {
	Relation map[string][]string `json:"datesLocations"`
}
