package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Vincent-Omondi/groupie-tracker/api"
	"github.com/Vincent-Omondi/groupie-tracker/controllers"
)

// Mock data for testing
var (
	mockArtists = []api.Artist{
		{ID: 1, Name: "Artist 1", Image: "image1.jpg", StartYear: 2000, FirstAlbum: "Album 1", Members: []string{"Member 1", "Member 2"}},
	}
	mockLocations = []api.Location{
		{ID: 1, Locations: []string{"Location 1", "Location 2"}},
	}
	mockDates = []api.Date{
		{ID: 1, Dates: []string{"2024-08-01", "2024-08-02"}},
	}
	mockRelations = []api.Relation{
		{ID: 1, DatesLocations: map[string][]string{"2024-08-01": {"Location 1"}, "2024-08-02": {"Location 2"}}},
	}
)

// Mock FetchData function
func mockFetchData(url string) ([]byte, error) {
	switch url {
	case api.ArtistsURL:
		return json.Marshal(mockArtists)
	case api.LocationsURL:
		return json.Marshal(mockLocations)
	case api.DatesURL:
		return json.Marshal(mockDates)
	case api.RelationURL:
		return json.Marshal(mockRelations)
	}
	return nil, nil
}

// TestGetArtistsHandler tests the GetArtistsHandler function
func TestGetArtistsHandler(t *testing.T) {
	// Replace FetchData with the mock implementation
	originalFetchData := api.FetchData
	defer func() { api.FetchData = originalFetchData }() // Restore original function after test

	api.FetchData = mockFetchData

	req, err := http.NewRequest("GET", "/artists", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetArtistsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var artists []api.Artist
	err = json.Unmarshal(rr.Body.Bytes(), &artists)
	if err != nil {
		t.Fatal(err)
	}

	if len(artists) != len(mockArtists) {
		t.Errorf("handler returned unexpected body: got %v want %v", len(artists), len(mockArtists))
	}
}

// TestGetLocationsHandler tests the GetLocationsHandler function
func TestGetLocationsHandler(t *testing.T) {
	// Replace FetchData with the mock implementation
	originalFetchData := api.FetchData
	defer func() { api.FetchData = originalFetchData }() // Restore original function after test

	api.FetchData = mockFetchData

	req, err := http.NewRequest("GET", "/locations", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetLocationsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var locations []api.Location
	err = json.Unmarshal(rr.Body.Bytes(), &locations)
	if err != nil {
		t.Fatal(err)
	}

	if len(locations) != len(mockLocations) {
		t.Errorf("handler returned unexpected body: got %v want %v", len(locations), len(mockLocations))
	}
}

// TestGetDatesHandler tests the GetDatesHandler function
func TestGetDatesHandler(t *testing.T) {
	// Replace FetchData with the mock implementation
	originalFetchData := api.FetchData
	defer func() { api.FetchData = originalFetchData }() // Restore original function after test

	api.FetchData = mockFetchData

	req, err := http.NewRequest("GET", "/dates", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetDatesHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var dates []api.Date
	err = json.Unmarshal(rr.Body.Bytes(), &dates)
	if err != nil {
		t.Fatal(err)
	}

	if len(dates) != len(mockDates) {
		t.Errorf("handler returned unexpected body: got %v want %v", len(dates), len(mockDates))
	}
}

// TestGetRelationsHandler tests the GetRelationsHandler function
func TestGetRelationsHandler(t *testing.T) {
	// Replace FetchData with the mock implementation
	originalFetchData := api.FetchData
	defer func() { api.FetchData = originalFetchData }() // Restore original function after test

	api.FetchData = mockFetchData

	req, err := http.NewRequest("GET", "/relations", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetRelationsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var relations []api.Relation
	err = json.Unmarshal(rr.Body.Bytes(), &relations)
	if err != nil {
		t.Fatal(err)
	}

	if len(relations) != len(mockRelations) {
		t.Errorf("handler returned unexpected body: got %v want %v", len(relations), len(mockRelations))
	}
}
