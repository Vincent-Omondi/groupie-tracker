//controllers/routes.go

package controllers

import (
	"net/http"
)

// Update RegisterRoutes function
func RegisterRoutes() {
	http.HandleFunc("/artists", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Has("id") {
			GetArtistByIDHandler(w, r)
		} else {
			GetArtistsHandler(w, r)
		}
	})
	http.HandleFunc("/locations", GetLocationsHandler)
	http.HandleFunc("/dates", GetDatesHandler)
	http.HandleFunc("/relations", GetRelationsHandler)
}

// // RegisterRoutes sets up the application routes
// func RegisterRoutes() {
// 	http.HandleFunc("/artists", GetArtistsHandler)
// 	http.HandleFunc("/locations", GetLocationsHandler)
// 	http.HandleFunc("/dates", GetDatesHandler)
// 	http.HandleFunc("/relations", GetRelationsHandler)

// 	// Serve static files (if any)
// 	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
// }
