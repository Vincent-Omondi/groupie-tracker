// controllers/handler.go
package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Vincent-Omondi/groupie-tracker/api"
)

// GetArtistsHandler handles the /artists route
func GetArtistsHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := api.GetArtists()
	if err != nil {
		http.Error(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(artists); err != nil {
		http.Error(w, "Failed to encode locations", http.StatusInternalServerError)
	}
}

// GetLocationsHandler handles the /locations route
func GetLocationsHandler(w http.ResponseWriter, r *http.Request) {
	locations, err := api.GetLocations()
	if err != nil {
		http.Error(w, "Failed to fetch locations", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(locations); err != nil {
		http.Error(w, "Failed to encode locations", http.StatusInternalServerError)
	}
}

// GetDatesHandler handles the /dates route
func GetDatesHandler(w http.ResponseWriter, r *http.Request) {
	dates, err := api.GetDates()
	if err != nil {
		http.Error(w, "Failed to fetch dates", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dates); err != nil {
		http.Error(w, "Failed to encode locations", http.StatusInternalServerError)
	}
}

// GetRelationsHandler handles the /relations route
func GetRelationsHandler(w http.ResponseWriter, r *http.Request) {
	relations, err := api.GetRelations()
	if err != nil {
		http.Error(w, "Failed to fetch relations", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(relations); err != nil {
		http.Error(w, "Failed to encode locations", http.StatusInternalServerError)
	}
}

// GetArtistByIDHandler handles the /artists?id={id} route
func GetArtistByIDHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Missing artist ID", http.StatusBadRequest)
		return
	}

	artistID, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	artists, err := api.GetArtists()
	if err != nil {
		http.Error(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}

	var selectedArtist api.Artist
	for _, artist := range artists {
		if artist.ID == artistID {
			selectedArtist = artist
			break
		}
	}

	if selectedArtist.ID == 0 {
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(selectedArtist); err != nil {
		http.Error(w, "Failed to encode artist", http.StatusInternalServerError)
	}
}
