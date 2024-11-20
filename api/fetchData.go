package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func fetchData(url string, data any) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(data)
}
