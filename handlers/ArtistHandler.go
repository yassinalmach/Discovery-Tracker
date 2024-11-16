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
	if ID, err := strconv.Atoi(id); err != nil || ID < 1 && ID > 52 {
		utils.PageNotFound(w, r)
		return
	}
	
	var data api.Artists
	api.FetchData(w, api.URL+"artists/"+id, &data)
	api.FetchData(w, api.URL+"locations/"+id, &data.Location)
	api.FetchData(w, api.URL+"dates/"+id, &data.Date)
	api.FetchData(w, api.URL+"relation/"+id, &data.Relation)
	
	parseTemplate(w, "artist.html", data)
}
