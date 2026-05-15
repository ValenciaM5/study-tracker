package main

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

// init runs once before tests, parsing templates so page handlers can render.
func init() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	w := httptest.NewRecorder()

	healthHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	expected := `{"status":"ok"}`
	if w.Body.String() != expected {
		t.Errorf("expected body %q, got %q", expected, w.Body.String())
	}
}

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	homeHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestUnknownPathReturns404(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/totally-not-a-page", nil)
	w := httptest.NewRecorder()

	homeHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}
}
