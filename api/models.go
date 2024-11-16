package api

type ArtistType struct {
	Id           int
	Name         string
	Image        string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Location     LocationType
	Date         DateType
	Relation     RelationType
}

type LocationType struct {
	Location []string
}

type DateType struct {
	Date []string
}
type RelationType struct {
	Relation map[string][]string
}
