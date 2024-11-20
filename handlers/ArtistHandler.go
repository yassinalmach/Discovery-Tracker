package handlers

import (
	"groupie/api"
	"groupie/utils"
	"net/http"
	"strconv"
	"strings"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/artist/")
	ID, err := strconv.Atoi(id)
	if err != nil || ID < 1 && ID > 52 {
		utils.PageNotFound(w, r)
		return
	}

	parseTemplate(w, "artist.html", api.AllArtists[ID-1])
}
