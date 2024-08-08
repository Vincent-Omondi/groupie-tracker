// backend/main.go

package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Vincent-Omondi/groupie-tracker/controllers"
)

// ServeTemplates renders the HTML templates
func ServeTemplates(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	controllers.RegisterRoutes()

	// Serve the main template at the root URL
	http.HandleFunc("/", ServeTemplates)

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
