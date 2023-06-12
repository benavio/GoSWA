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
		t.Fatal("statos not OK")
	}
}

func TestAlbumDetail(t *testing.T) {
	IDalbum := "1"
	request, _ := http.NewRequest("GET", "/albums/"+IDalbum, strings.NewReader(""))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusOK {
		t.Fatal("statos not OK")
	}
}

func TestAlbumNotFound(t *testing.T) {
	IDalbum := "14"
	request, _ := http.NewRequest("GET", "/albums/"+IDalbum, strings.NewReader(""))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusNotFound {
		t.Fatal("status 404")
	}
}

func TestDeleteAlbum(t *testing.T) {
	IDalbum := "1"
	request, _ := http.NewRequest("DELETE", "/albums/"+IDalbum, strings.NewReader(""))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusNoContent {
		t.Fatal("status 204")
	}

}
func TestDeleteAlbumNotFound(t *testing.T) {
	IDalbum := "1"
	request, _ := http.NewRequest("DELETE", "/albums/"+IDalbum, strings.NewReader(""))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusNotFound {
		t.Fatal("status 404")
	}
}

func TestUpdateAlbum(t *testing.T) {
	IDalbum := "2"
	request, _ := http.NewRequest("PUT", "/albums/"+IDalbum, strings.NewReader(`{"title": "Karlos Makaroni"}`))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusOK {
		t.Fatal("status 200")
	}
}

func TestUpdateAlbumNotFound(t *testing.T) {
	IDalbum := "999"
	request, _ := http.NewRequest("PUT", "/albums/"+IDalbum, strings.NewReader(""))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusNotFound {
		t.Fatal("status 404")
	}
}

func TestUpdateBadStructure(t *testing.T) {
	IDalbum := "2"
	request, _ := http.NewRequest("PUT", "/albums/"+IDalbum, strings.NewReader(`{"title": "Karlos Makaroni"}`))
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
		t.Fatal(w.Code)
	}
}

func TestCreateAlbums(t *testing.T) {
	request, _ := http.NewRequest("POST", "/albums", strings.NewReader(`{"title": "Karlos Makaroni"}`))
	w := httptest.NewRecorder()
	handleRequest(w, request)
	if w.Code != http.StatusBadRequest {
		t.Fatal("status created", w.Code)
	}
}
