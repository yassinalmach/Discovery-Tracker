package handlers

import (
	"fmt"
	"groupie/api"
	"groupie/features"
	"groupie/utils"
	"net/http"
	"strconv"
	"strings"
)

type artistData struct {
	Artist     api.Artists
	Coordinate []features.CoordinateModel
}

// ArtistHandler handles requests for individual artist pages
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the artist ID from the requested URL
	id := strings.TrimPrefix(r.URL.Path, "/artist/")
	ID, err := strconv.Atoi(id)
	if err != nil || ID < 1 && ID > 52 {
		utils.PageNotFound(w, r)
		return
	}
	
	// add map to artist page
	var coordinate []features.CoordinateModel
	for _, loc := range api.AllLocations.Index[ID-1].Location {
		coord, err := features.GetCoordinates(loc)
		if err != nil {
			fmt.Printf("Error geocode %s: %v\n", loc, err)
			continue
		}

		lat := strconv.FormatFloat(coord.Result[0].Geometry.Location.Lat, 'f', -1, 64)
		lon := strconv.FormatFloat(coord.Result[0].Geometry.Location.Lon, 'f', -1, 64)

		coordinate = append(coordinate, features.CoordinateModel{Location: loc, Addr: lat + "," + lon})
	}

	pageData := artistData{
		Artist:     api.AllArtists[ID-1],
		Coordinate: coordinate,
	}

	parseTemplate(w, "artist.html", pageData)
}
