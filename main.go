// backend/main.go

package main

import (
	"log"
	"net/http"

	"github.com/Vincent-Omondi/groupie-tracker/controllers"
)

func main() {
	controllers.RegisterRoutes()
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
