package handlers

import (
	"net/http"
	"os"
	"path/filepath"
)

// ServeStaticFiles returns a handler function for serving static files
func ServeStaticFiles(w http.ResponseWriter, r *http.Request) {
	// Get clean file path and prevent directory traversal
	path := filepath.Clean("." + r.URL.Path)
	if info, err := os.Stat(path); err != nil || info.IsDir() {
		http.Error(w, "internal server error", 500)
	}

	http.ServeFile(w, r, path)
}
