// routes.go

package controllers

import (
	"net/http"
)

// RegisterRoutes sets up the application routes
func RegisterRoutes() {
	http.HandleFunc("/artists", GetArtistsHandler)
	http.HandleFunc("/locations", GetLocationsHandler)
	http.HandleFunc("/dates", GetDatesHandler)
	http.HandleFunc("/relations", GetRelationsHandler)

	// Serve static files (if any)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}


