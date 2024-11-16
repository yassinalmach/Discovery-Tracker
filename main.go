package main

import (
	"fmt"
	"groupie/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	http.HandleFunc("GET /", handlers.HomeHandler)
	http.HandleFunc("GET /artist/", handlers.ArtistHandler)
	http.HandleFunc("GET /static/", handlers.ServeStaticFiles)

	http.ListenAndServe(":8080", mux)
	fmt.Println("server running on: http://localhost:8080")
}
