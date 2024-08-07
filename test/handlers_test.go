package tests

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Vincent-Omondi/groupie-tracker/api"
	"github.com/Vincent-Omondi/groupie-tracker/controllers"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/assert"
)

type mockFetcher struct {
	data map[string][]byte
	err  error
}

func (m mockFetcher) FetchData(url string) ([]byte, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.data[url], nil
}

func TestGetArtistsHandler(t *testing.T) {
	data := map[string][]byte{
		api.ArtistsURL: []byte(`[{"id":1,"name":"Artist1","image":"image1","creationDate":2000,"firstAlbum":"2001-01-01","members":["member1"],"locations":"loc1","concertDates":"date1","relations":"rel1"}]`),
	}

	api.Fetcher = mockFetcher{data: data}

	req, err := http.NewRequest("GET", "/artists", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetArtistsHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var artists []api.Artist
	err = json.NewDecoder(rr.Body).Decode(&artists)
	assert.NoError(t, err)
	assert.Len(t, artists, 1)
	assert.Equal(t, "Artist1", artists[0].Name)
}

func TestGetArtistsHandler_Failure(t *testing.T) {
	api.Fetcher = mockFetcher{err: errors.New("fetch error")}

	req, err := http.NewRequest("GET", "/artists", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetArtistsHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "Failed to fetch artists\n", rr.Body.String())
}

func TestGetLocationsHandler(t *testing.T) {
	data := map[string][]byte{
		api.LocationsURL: []byte(`{"index":[{"id":1,"locations":["loc1"],"dates":"date1"}]}`),
	}

	api.Fetcher = mockFetcher{data: data}

	req, err := http.NewRequest("GET", "/locations", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetLocationsHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var locations []api.Location
	err = json.NewDecoder(rr.Body).Decode(&locations)
	assert.NoError(t, err)
	assert.Len(t, locations, 1)
	assert.Equal(t, "loc1", locations[0].Locations[0])
}

func TestGetLocationsHandler_Failure(t *testing.T) {
	api.Fetcher = mockFetcher{err: errors.New("fetch error")}

	req, err := http.NewRequest("GET", "/locations", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetLocationsHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "Failed to fetch locations\n", rr.Body.String())
}

func TestGetDatesHandler(t *testing.T) {
	data := map[string][]byte{
		api.DatesURL: []byte(`{"index":[{"id":1,"dates":["date1"]}]}`),
	}

	api.Fetcher = mockFetcher{data: data}

	req, err := http.NewRequest("GET", "/dates", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetDatesHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var dates []api.Date
	err = json.NewDecoder(rr.Body).Decode(&dates)
	assert.NoError(t, err)
	assert.Len(t, dates, 1)
	assert.Equal(t, "date1", dates[0].Dates[0])
}

func TestGetDatesHandler_Failure(t *testing.T) {
	api.Fetcher = mockFetcher{err: errors.New("fetch error")}

	req, err := http.NewRequest("GET", "/dates", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetDatesHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "Failed to fetch dates\n", rr.Body.String())
}

func TestGetRelationsHandler(t *testing.T) {
	data := map[string][]byte{
		api.RelationURL: []byte(`{"index":[{"id":1,"datesLocations":{"loc1":["date1"]}}]}`),
	}

	api.Fetcher = mockFetcher{data: data}

	req, err := http.NewRequest("GET", "/relations", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetRelationsHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var relations []api.Relation
	err = json.NewDecoder(rr.Body).Decode(&relations)
	assert.NoError(t, err)
	assert.Len(t, relations, 1)
	assert.Equal(t, "date1", relations[0].DatesLocations["date1"][0])
}

func TestGetRelationsHandler_Failure(t *testing.T) {
	api.Fetcher = mockFetcher{err: errors.New("fetch error")}

	req, err := http.NewRequest("GET", "/relations", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetRelationsHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "Failed to fetch relations\n", rr.Body.String())
}
