package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUploadHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodPost, "http://localhost/upload", nil)
	uploadHandler(w, r)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected %d, get %d", http.StatusCreated, w.Code)
	}
}

func TestResultHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "http://localhost/artifact/hash", nil)
	artifactHandler(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Expected %d, get %d", http.StatusOK, w.Code)
	}
}

func BenchmarkUploadHander(b *testing.B) {

	for i := 0; i < b.N; i++ {

		req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/upload", nil)

		rec := httptest.NewRecorder()

		uploadHandler(rec, req)

		if rec.Code != http.StatusCreated {
			b.Errorf("Expected %d, get %d", http.StatusCreated, rec.Code)
		}
	}
}

func BenchmarkResultHander(b *testing.B) {

	for i := 0; i < b.N; i++ {

		req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/artifact/hash", nil)

		rec := httptest.NewRecorder()

		artifactHandler(rec, req)

		if rec.Code != http.StatusOK {
			b.Errorf("Expected %d, get %d", http.StatusOK, rec.Code)
		}

		if rec.Code != http.StatusOK {
			b.Errorf("Expected %d, get %d", http.StatusOK, rec.Code)
		}
	}
}
