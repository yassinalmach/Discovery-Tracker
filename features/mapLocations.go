package features

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CoordinateModel struct {
	Location string
	Addr     string
}

type ApiModel struct {
	Result []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
}

func GetCoordinates(location string) (ApiModel, error) {
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=API_KEY", location)

	resp, err := http.Get(url)
	if err != nil {
		return ApiModel{}, err
	}
	defer resp.Body.Close()

	var coord ApiModel
	if err := json.NewDecoder(resp.Body).Decode(&coord); err != nil {
		return ApiModel{}, err
	}

	return coord, nil
}
