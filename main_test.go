package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAlbumsList(t *testing.T) {
	router := getRouter()
	request, _ := http.NewRequest("GET", "/albums", strings.NewReader(""))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	if w.Code != http.StatusOK {
		t.Fatal("statos not OK")
	}
}

func TestAlbumDetail(t *testing.T) {
	router := getRouter()
	IDalbum := "1"
	request, _ := http.NewRequest("GET", "/albums/"+IDalbum, strings.NewReader(""))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	if w.Code != http.StatusOK {
		t.Fatal("statos not OK")
	}
}

func TestNotFound(t *testing.T) {
	router := getRouter()
	IDalbum := "14"
	request, _ := http.NewRequest("GET", "/albums/"+IDalbum, strings.NewReader(""))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	if w.Code != http.StatusNotFound {
		t.Fatal("status 404")
	}
}
