package main

import (
	"fmt"
	"groupie/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("GET /artist/", handlers.ArtistHandler)
	mux.HandleFunc("GET /static/", handlers.ServeStaticFiles)

	fmt.Println("server running on: http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
