// api/api.go
package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Structs to unmarshal JSON data
type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Members      []string `json:"members"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

const (
	ArtistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	LocationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	DatesURL     = "https://groupietrackers.herokuapp.com/api/dates"
	RelationURL  = "https://groupietrackers.herokuapp.com/api/relation"
)

// DataFetcher interface to abstract the data fetching logic
type DataFetcher interface {
	FetchData(url string) ([]byte, error)
}

// Default Fetcher implementation with set timeout
var Fetcher DataFetcher = fetcher{
	client: &http.Client{
		Timeout: 20 * time.Second,
	},
}

// fetcher struct that implements the DataFetcher interface
type fetcher struct {
	client *http.Client
}

// FetchData makes an HTTP GET request to the given URL and returns the response body
func (f fetcher) FetchData(url string) ([]byte, error) {
	resp, err := f.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return body, nil
}

// GetArtists fetches the artist data from the API and returns a slice of Artist structs
func GetArtists() ([]Artist, error) {
	body, err := Fetcher.FetchData(ArtistsURL)
	if err != nil {
		return nil, err
	}

	var artists []Artist
	if err := json.Unmarshal(body, &artists); err != nil {
		return nil, fmt.Errorf("failed to unmarshal artists: %v", err)
	}

	return artists, nil
}

// GetLocations fetches the location data from the API and returns a slice of Location structs
func GetLocations() ([]Location, error) {
	body, err := Fetcher.FetchData(LocationsURL)
	if err != nil {
		return nil, err
	}

	var locations struct {
		Index []Location `json:"index"`
	}

	if err := json.Unmarshal(body, &locations); err != nil {
		return nil, fmt.Errorf("failed to unmarshal locations: %v", err)
	}

	return locations.Index, nil
}

// GetDates fetches the date data from the API and returns a slice of Date structs
func GetDates() ([]Date, error) {
	body, err := Fetcher.FetchData(DatesURL)
	if err != nil {
		return nil, err
	}

	var dates struct {
		Index []Date `json:"index"`
	}

	if err := json.Unmarshal(body, &dates); err != nil {
		return nil, fmt.Errorf("failed to unmarshal dates: %v", err)
	}

	return dates.Index, nil
}

// GetRelations fetches the relation data from the API and returns a slice of Relation structs
func GetRelations() ([]Relation, error) {
	body, err := Fetcher.FetchData(RelationURL)
	if err != nil {
		return nil, err
	}

	var relations struct {
		Index []Relation `json:"index"`
	}
	if err := json.Unmarshal(body, &relations); err != nil {
		return nil, fmt.Errorf("failed to unmarshal relations: %v", err)
	}

	return relations.Index, nil
}
