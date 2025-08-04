package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func addAlbumHandler(resp http.ResponseWriter, req *http.Request) {
	log.Debug("addAlbumHandler")
}

func addReviewHandler(resp http.ResponseWriter, req *http.Request) {
	log.Debug("addReviewHandler")
}
