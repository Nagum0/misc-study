package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type albumRequest struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
}

// Get all of the albums from the database
func getAllAlbums(resp http.ResponseWriter, req *http.Request) {
	log.Info("[getAllAlbums] Request at ", req.URL.Path)

	if req.Method != http.MethodGet {
		log.Error("[getAllAlbums] Error: Expected GET request")
		http.Error(resp, "[getAllAlbums] Error: Expected GET requestn\n", http.StatusBadRequest)
		return
	}

	var albums []Album
	err = db.Select(&albums, "SELECT * FROM Albums")
	if err != nil {
		log.Error("[getAllAlbums] Error: ", err.Error())
		http.Error(resp, fmt.Sprintf("[getAllAlbums] Error: %v\n", err.Error()), http.StatusInternalServerError)
		return
	}

	printAlbums(resp, albums)
}

// Get a specific album from the database
func getAlbum(resp http.ResponseWriter, req *http.Request) {
	log.Info("[getAlbum] Request at ", req.URL.Path)

	if req.Method != http.MethodPost {
		log.Error("[getAlbum] Error: Expected POST request")
		http.Error(resp, "", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Error("[getAlbum] Error: ", err.Error())
		http.Error(resp, fmt.Sprintf("[getAlbum] Error: %v\n", err.Error()), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	var albumRequest albumRequest
	err = json.Unmarshal(body, &albumRequest)
	if err != nil {
		log.Error("[getAlbum] Error: ", err.Error())
		http.Error(resp, fmt.Sprintf("[getAlbum] Error: %v\n", err.Error()), http.StatusBadRequest)
		return
	}

	if albumRequest.Name == "" {
		log.Error("[getAlbum] Error: Album name cannot be empty")
		http.Error(resp, "[getAlbum] Error: Album name cannot be empty", http.StatusBadRequest)
		return
	}

	var albums []Album
	if albumRequest.Artist == "" {
		albums, err = selectAlbumsByName(albumRequest.Name)
	} else {
		albums, err = selectAlbumsByArtistAndName(albumRequest.Name, albumRequest.Artist)
	}

	if err != nil {
		log.Error("[getAlbum] Error: ", err.Error())
		http.Error(resp, fmt.Sprintf("[getAlbum] Error: %v\n", err.Error()), http.StatusInternalServerError)
		return
	}

	printAlbums(resp, albums)
}

func addAlbumHandler(resp http.ResponseWriter, req *http.Request) {
	log.Debug("addAlbumHandler -- NOT IMPLEMENTED")
	fmt.Fprintf(resp, "addAlbumHandler -- NOT IMPLEMENTED")
}

func addReviewHandler(resp http.ResponseWriter, req *http.Request) {
	log.Debug("addReviewHandler -- NOT IMPLEMENTED")
	fmt.Fprintf(resp, "addReviewHandler -- NOT IMPLEMENTED")
}
