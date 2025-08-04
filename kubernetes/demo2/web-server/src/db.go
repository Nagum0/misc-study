package main

import (
	"fmt"
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
