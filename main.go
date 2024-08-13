package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"learn.zone01kisumu.ke/git/johnodhiambo0/groupie-tracker/api"
)

type TemplateData struct {
	Artists []api.Artist
}

type ArtistDetailData struct {
	Artist   api.Artist
	Relation api.Relation
}

var artistCache []api.Artist
var cacheTime time.Time

func ServeArtists(w http.ResponseWriter, r *http.Request) {
	if time.Since(cacheTime) > time.Minute*10{
		artists, err := api.GetArtists()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		artistCache = artists
		cacheTime = time.Now()
	}

	data := TemplateData{Artists: artistCache}

	tmpl := template.Must(template.ParseFiles("templates/artists.html"))
	tmpl.Execute(w, data)
}

func ServeArtistDetails(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/artist/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	artist, relation, err := api.GetArtistByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := ArtistDetailData{
		Artist:   *artist,
		Relation: *relation,
	}

	tmpl := template.Must(template.ParseFiles("templates/artist_details.html"))
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing template artist_details:", err)
	}
}

func main() {
	http.HandleFunc("/", ServeArtists)
	http.HandleFunc("/artist/", ServeArtistDetails)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
