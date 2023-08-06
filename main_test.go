package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func handleRequest(w *httptest.ResponseRecorder, r *http.Request) {
	router := getRouter()
	router.ServeHTTP(w, r)
}

func TestAlbumsList(t *testing.T) {
	request, _ := http.NewRequest("GET", "/albums", strings.NewReader(""))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusOK {
		t.Fatal("status not OK")
	}
}

func TestAlbumDetail(t *testing.T) {
	albumID := "3"
	request, _ := http.NewRequest("GET", "/albums/"+albumID, strings.NewReader(""))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusOK {
		t.Fatal("status not OK", w.Code)
	}
}

func TestAlbumNotFound(t *testing.T) {
	albumID := "9999"
	request, _ := http.NewRequest("GET", "/albums/"+albumID, strings.NewReader(""))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusNotFound {
		t.Fatal("status 404", w.Code)
	}
}

func TestDeleteAlbum(t *testing.T) {
	albumID := "1"
	request, _ := http.NewRequest("DELETE", "/albums/"+albumID, strings.NewReader(""))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusNoContent {
		t.Fatal("status 204")
	}

}
func TestDeleteAlbumNotFound(t *testing.T) {
	albumID := "999999"
	request, _ := http.NewRequest("DELETE", "/albums/"+albumID, strings.NewReader(""))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusNotFound {
		t.Fatal("status 404")
	}
}

func TestUpdateAlbum(t *testing.T) {
	albumID := "3"
	request, _ := http.NewRequest("PUT", "/albums/"+albumID, strings.NewReader(`{"id": "4", "title": "TEST", "artist": "TEST", "price": 56.99}`))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusOK {
		t.Fatal("status 200")
	}
}

func TestUpdateAlbumNotFound(t *testing.T) {
	albumID := "99999999999"
	request, _ := http.NewRequest("PUT", "/albums/"+albumID, strings.NewReader(`{"id": "4", "title": "Gib Beam", "artist": "John Coltrane", "price": 56.99}`))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusNotFound {
		t.Fatal("status 404")
	}
}

func TestUpdateBadStructure(t *testing.T) {
	albumID := "9999"
	request, _ := http.NewRequest("POST", "/albums/"+albumID, strings.NewReader(`{"title": "Karlos Makaroni"}`))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusNotFound {
		t.Fatal("status 404")
	}
}

func TestCreateBadStructure(t *testing.T) {
	request, _ := http.NewRequest("POST", "/albums", strings.NewReader(""))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusBadRequest {
		t.Fatal("status 400", w.Code)
	}
}

func TestCreateAlbums(t *testing.T) {
	request, _ := http.NewRequest("POST", "/albums", strings.NewReader(`{"id": "4", "title": "Gib Beam", "artist": "John Coltrane", "price": 56.99}`))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusCreated {
		t.Fatal("status created 201", w.Code)
	}
}
