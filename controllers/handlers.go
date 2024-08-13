package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"learn.zone01kisumu.ke/git/johnodhiambo0/groupie-tracker/api"
)

type TemplateData struct {
	Artists   []api.Artist
	Query     string
	NoResults bool
}

type ArtistDetailData struct {
	Artist   api.Artist
	Relation struct {
		Locations      []string
		Dates          []string
		DatesLocations map[string][]string
	}
}

func ErrorHandler(w http.ResponseWriter, message string, statusCode int) {
	// Set the status code
	w.WriteHeader(statusCode)

	// Define error template data
	data := struct {
		StatusCode int
		ErrMsg     string
	}{
		StatusCode: statusCode,
		ErrMsg:     message,
	}

	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Error executing error template")
		http.Error(w, "Error executing data deatils", http.StatusInternalServerError)
		return
	}
}

func ServeArtists(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	artists, err := api.GetArtists()
	if err != nil {
		log.Printf("Error getting artists: %v", err)
		ErrorHandler(w, "Unable to retrieve artists at this time. Please try again later.", http.StatusInternalServerError)
		return
	}

	filteredArtists := filterArtists(artists, query)

	// Check if no results were found and query is not empty
	if len(filteredArtists) == 0 && query != "" {
		ErrorHandler(w, "We couldn't find any artists matching your search criteria. Please try a different term or check your spelling.", http.StatusNotFound)
		return
	}

	data := TemplateData{
		Artists:   filteredArtists,
		Query:     query,
		NoResults: len(filteredArtists) == 0 && query != "",
	}

	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		ErrorHandler(w, "An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		ErrorHandler(w, "We encountered an issue while rendering the page. Please try again later.", http.StatusInternalServerError)
		return
	}
}

// filterArtists filters the list of artists based on the search query
func filterArtists(artists []api.Artist, query string) []api.Artist {
	if query == "" {
		return artists
	}

	var result []api.Artist
	for _, a := range artists {
		if strings.Contains(strings.ToLower(a.Name), strings.ToLower(query)) {
			result = append(result, a)
		}
	}
	return result
}

func ServeArtistDetails(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/artist/" {
		ErrorHandler(w, "page not found", http.StatusNotExtended)
		return
	}

	idStr := r.URL.Path[len("/artist/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorHandler(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	artist, relation, err := api.GetArtistByID(id)
	if err != nil {
		log.Printf("Error retrieving artist by ID %v: %s", id, err)
		ErrorHandler(w, "Unable to retrieve artist details at this time. Please try again later.", http.StatusInternalServerError)
		return
	}

	locations := make([]string, 0, len(relation.DatesLocations))
	dates := make([]string, 0)
	for location, datelist := range relation.DatesLocations {
		locations = append(locations, location)
		dates = append(dates, datelist...)
	}

	data := ArtistDetailData{
		Artist: *artist,
		Relation: struct {
			Locations      []string
			Dates          []string
			DatesLocations map[string][]string
		}{
			Locations:      locations,
			Dates:          dates,
			DatesLocations: relation.DatesLocations,
		},
	}

	tmpl, err := template.ParseFiles("templates/artist_details.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		ErrorHandler(w, "Unable to load artist details at this time. Please try again later.", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		ErrorHandler(w, "Error rendering artist details. Please try again later.", http.StatusInternalServerError)
		return
	}
}

// GetArtistsHandler handles the /artists route
func GetArtistsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	artists, err := api.GetArtists()
	if err != nil {
		log.Printf("Error fetching artists: %v", err)
		ErrorHandler(w, "Unable to retrieve artist information at this time. Please try again later.", http.StatusInternalServerError)
		return
	}

	filteredArtists := filterArtists(artists, query)

	if len(filteredArtists) == 0 {
		ErrorHandler(w, "No artists found matching the search term.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(filteredArtists); err != nil {
		log.Printf("Error encoding artists data to JSON: %v", err)
		ErrorHandler(w, "An error occurred while processing the artist data. Please try again later.", http.StatusInternalServerError)
		return
	}
}

// GetLocationsHandler handles the /locations route
func GetLocationsHandler(w http.ResponseWriter, r *http.Request) {
	locations, err := api.GetLocations()
	if err != nil {
		log.Printf("Error fetching locations: %v", err)
		ErrorHandler(w, "Unable to retrieve locations at this time. Please try again later.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(locations); err != nil {
		log.Printf("Error encoding locations data to JSON: %v", err)
		ErrorHandler(w, "An error occurred while processing location data. Please try again later.", http.StatusInternalServerError)
		return
	}
}

// GetDatesHandler handles the /dates route
func GetDatesHandler(w http.ResponseWriter, r *http.Request) {
	dates, err := api.GetDates()
	if err != nil {
		log.Printf("Error fetching dates: %v", err)
		ErrorHandler(w, "Unable to retrieve dates at this time. Please try again later.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dates); err != nil {
		log.Printf("Error encoding dates data to JSON: %v", err)
		ErrorHandler(w, "An error occurred while processing date information. Please try again later.", http.StatusInternalServerError)
		return
	}
}

// GetRelationsHandler handles the /relations route
func GetRelationsHandler(w http.ResponseWriter, r *http.Request) {
	relations, err := api.GetRelations()
	if err != nil {
		log.Printf("Error fetching relations: %v", err)
		ErrorHandler(w, "Unable to retrieve relations at this time. Please try again later.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(relations); err != nil {
		log.Printf("Error encoding relations data to JSON: %v", err)
		ErrorHandler(w, "An error occurred while processing relation data. Please try again later.", http.StatusInternalServerError)
		return
	}
}

func GetArtistByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Get artist ID from the URL path
	idStr := r.URL.Path[len("/artists/"):]
	artistID, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Invalid artist ID: %v", err)
		ErrorHandler(w, "Invalid artist ID provided. Please check and try again.", http.StatusBadRequest)
		return
	}

	artist, relation, err := api.GetArtistByID(artistID)
	if err != nil {
		// Check if the artist was not found
		if err.Error() == "artist not found" {
			ErrorHandler(w, "Artist not found. Please check the ID and try again.", http.StatusNotFound)
		} else {
			log.Printf("Error fetching artist or relation with ID %d: %v", artistID, err)
			ErrorHandler(w, "Unable to retrieve artist details at this time. Please try again later.", http.StatusInternalServerError)
		}
		return
	}

	// Create a response combining artist and relation data
	response := struct {
		Artist   *api.Artist   `json:"artist"`
		Relation *api.Relation `json:"relation"`
	}{
		Artist:   artist,
		Relation: relation,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response for artist ID %d: %v", artistID, err)
		ErrorHandler(w, "An error occurred while processing the response. Please try again later.", http.StatusInternalServerError)
		return
	}
}
