package main

import (
	"fmt"
	"io"
)

type Album struct {
	Id          int    `db:"id"`
	Artist      string `db:"artist"`
	Name        string `db:"name"`
	ReleaseDate string `db:"release_date"`
}

type Review struct {
	Id      int    `db:"id"`
	AlbumId int    `db:"album_id"`
	User    string `db:"user"`
	Review  string `db:"review"`
}

func getDSN() string {
	return fmt.Sprintf(
		"%v:%v@(%v:%v)/%v",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
}

func selectAlbumsByArtistAndName(name string, artist string) ([]Album, error) {
	var albums []Album
	err = db.Select(&albums, "SELECT * FROM Albums WHERE name = ? AND artist = ?", name, artist)
	return albums, err
}

func selectAlbumsByName(name string) ([]Album, error) {
	var albums []Album
	err = db.Select(&albums, "SELECT * FROM Albums WHERE name = ?", name)
	return albums, err
}

func printAlbums(w io.Writer, albums []Album) {
	for _, album := range albums {
		fmt.Fprintf(
			w,
			"ID: %v  | NAME: %v  | ARTIST: %v  | RELEASE DATE: %v\n",
			album.Id, album.Name, album.Artist, album.ReleaseDate,
		)
	}
}
