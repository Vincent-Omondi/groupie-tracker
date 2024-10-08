package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

var (
	ArtistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	LocationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	DatesURL     = "https://groupietrackers.herokuapp.com/api/dates"
	RelationURL  = "https://groupietrackers.herokuapp.com/api/relation"
)

// FetchData makes an HTTP GET request to the given URL and returns the response body
func FetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)
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
	body, err := FetchData(ArtistsURL)
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
	body, err := FetchData(LocationsURL)
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
	body, err := FetchData(DatesURL)
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
	body, err := FetchData(RelationURL)
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

// GetArtistByID fetches the artist data by ID and returns the Artist struct along with its relation
func GetArtistByID(artistID int) (*Artist, *Location, *Date, *Relation, error) {
    // Fetch artist data
    artists, err := GetArtists()
    if err != nil {
        return nil, nil, nil, nil, err
    }

    // Find the artist with the specified ID
    var artist *Artist
    for _, a := range artists {
        if a.ID == artistID {
            artist = &a
            break
        }
    }
    if artist == nil {
        return nil, nil, nil, nil, fmt.Errorf("artist not found")
    }

    // Fetch location data
    locations, err := GetLocations()
    if err != nil {
        return nil, nil, nil, nil, err
    }

    // Find the location for the specific artist
    var location *Location
    for _, l := range locations {
        if l.ID == artistID {
            location = &l
            break
        }
    }
    if location == nil {
        return nil, nil, nil, nil, fmt.Errorf("location not found for artist")
    }

    // Fetch date data
    dates, err := GetDates()
    if err != nil {
        return nil, nil, nil, nil, err
    }

    // Find the date for the specific artist
    var date *Date
    for _, d := range dates {
        if d.ID == artistID {
            date = &d
            break
        }
    }
    if date == nil {
        return nil, nil, nil, nil, fmt.Errorf("date not found for artist")
    }

    // Fetch relation data
    relations, err := GetRelations()
    if err != nil {
        return nil, nil, nil, nil, err
    }

    // Find the relation for the specific artist
    var relation *Relation
    for _, r := range relations {
        if r.ID == artistID {
            relation = &r
            break
        }
    }
    if relation == nil {
        return nil, nil, nil, nil, fmt.Errorf("relation not found for artist")
    }

    return artist, location, date, relation, nil
}
