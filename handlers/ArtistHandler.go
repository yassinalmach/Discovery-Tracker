package handlers

import (
	"groupie/utils"
	"net/http"
	"strconv"
	"strings"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/artist/")
	if ID, err := strconv.Atoi(id); err != nil || ID < 1 && ID > 52 {
		utils.PageNotFound(w, r)
		return
	}

	// parseTemplate(w, "artist.html", data)
}
