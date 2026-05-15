package main

import (
	"html/template"
	"log"
	"net/http"
)

// templates holds all parsed HTML templates from the templates/ folder.
// We declare it at package level so handlers and tests can reach it.
var templates *template.Template

func main() {
	// Parse every .html file in the templates folder.
	// template.Must panics if parsing fails — which is what we want at startup.
	templates = template.Must(template.ParseGlob("templates/*.html"))

	// Serve files in /static (CSS, images, etc.) under the /static/ URL prefix.
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Page routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/courses", coursesHandler)
	http.HandleFunc("/goals", goalsHandler)
	http.HandleFunc("/habits", habitsHandler)
	http.HandleFunc("/about", aboutHandler)

	// Health check route — useful for Docker, Kubernetes, and load balancers later.
	http.HandleFunc("/healthz", healthHandler)

	port := ":8080"
	log.Printf("Study Tracker is running on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// render is a small helper so each handler stays short and readable.
func render(w http.ResponseWriter, page string, data interface{}) {
	if err := templates.ExecuteTemplate(w, page, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// http.HandleFunc("/", ...) matches everything that isn't matched elsewhere,
	// so guard against unknown paths landing on the home page.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	render(w, "home.html", nil)
}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "courses.html", nil)
}

func goalsHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "goals.html", nil)
}

func habitsHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "habits.html", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "about.html", nil)
}

// healthHandler returns a tiny JSON payload so Kubernetes / Docker / monitors
// can quickly check the app is alive.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}
