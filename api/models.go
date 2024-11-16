package api

const URL = "https://groupietrackers.herokuapp.com/api"

type Artists struct {
	Id           int
	Name         string
	Image        string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Location     Locations
	Date         Dates
	Relation     Relations
}

type Locations struct {
	Location []string
}

type Dates struct {
	Date []string
}
type Relations struct {
	Relation map[string][]string
}
