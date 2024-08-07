package tests

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/Vincent-Omondi/groupie-tracker/controllers"
// )

// func TestRegisterRoutes(t *testing.T) {
// 	controllers.RegisterRoutes()

// 	routes := []string{
// 		"/artists",
// 		"/locations",
// 		"/dates",
// 		"/relations",
// 	}

// 	for _, route := range routes {
// 		req, err := http.NewRequest("GET", route, nil)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		rr := httptest.NewRecorder()
// 		http.DefaultServeMux.ServeHTTP(rr, req)

// 		if status := rr.Code; status != http.StatusOK {
// 			t.Errorf("route %v returned wrong status code: got %v want %v", route, status, http.StatusOK)
// 		}
// 	}
// }
