package api

import "sync"

var (
	AllData      []Data
	AllArtists   []Artists
	AllLocations LocationsData
	AllDates     DatesData
	AllRelations RelationsData
)

func init() {
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		fetchData(URL+"artists", &AllArtists)
	}()

	go func() {
		defer wg.Done()
		fetchData(URL+"locations", &AllLocations)
	}()

	go func() {
		defer wg.Done()
		fetchData(URL+"dates", &AllDates)
	}()

	go func() {
		defer wg.Done()
		fetchData(URL+"relations", &AllRelations)
	}()
	wg.Wait()

	for i := 0; i < 52; i++ {
		data := Data{
			Artist:   AllArtists[i],
			Location: AllLocations.Index[i],
			Date:     AllDates.Index[i],
			Relation: AllRelations.Index[i],
		}
		AllData = append(AllData, data)
	}
}
