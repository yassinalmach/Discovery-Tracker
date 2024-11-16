package api

import (
	"encoding/json"
	"net/http"
)

func FetchData(w http.ResponseWriter, url string, data any) {
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "internal server error", 500)
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(data)
}
