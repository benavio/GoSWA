package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Storage interface {
	Create() album
	Read() album
	ReadID() (album, error)
	Update() album
	Delete() album
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
type HttpError struct {
	Error string `json:"error"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jery", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan", Artist: "Sarah Vaughan", Price: 39.99},
}

type MemoryStorage struct {
	albums []album
}

var storage = NewMemoryStorage()

func NewMemoryStorage() MemoryStorage {
	var albums = []album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jery", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan", Artist: "Sarah Vaughan", Price: 39.99},
	}
	return MemoryStorage{albums: albums}
}

func (s MemoryStorage) Create(am album) album {
	s.albums = append(s.albums, am)
	return am
}
func (s MemoryStorage) ReadID(id string) (album, error) {
	for _, a := range albums {
		if a.ID == id {
			return a, nil
		}
	}
	return album{}, errors.New("not found")
}

func (s MemoryStorage) Read() []album {
	return s.albums
}

func (s MemoryStorage) Update(id string, newAlbum album) (album, error) {
	for i := range s.albums {
		if s.albums[i].ID == id {
			s.albums[i] = newAlbum
			return s.albums[i], nil
		}
	}
	return album{}, errors.New("not found")
}

func (s MemoryStorage) Delete(id string) error {
	for i, a := range s.albums {
		if a.ID == id {
			s.albums = append(s.albums[:i], s.albums[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

func getAlbum(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, storage.Read())
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, HttpError{"bad_request"})
		return
	}
	storage.Create(newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	album, err := storage.ReadID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, HttpError{"not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func updateAlbumsById(c *gin.Context) {
	id := c.Param("id")
	var newAlbum album
	c.BindJSON(&newAlbum)
	album, err := storage.Update(id, newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, HttpError{"not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func deleteAlbumById(c *gin.Context) {
	id := c.Param("id")
	err := storage.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, HttpError{"not found"})
	}
	c.IndentedJSON(http.StatusNoContent, album{})

}

func getRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/albums", getAlbum)
	router.GET("/albums/:id", getAlbumById)
	router.DELETE("/albums/:id", deleteAlbumById)
	router.PUT("/albums/:id", updateAlbumsById)
	router.POST("/albums", postAlbums)
	return router
}

func main() {
	router := getRouter()
	router.Run("localhost:8080")
}
