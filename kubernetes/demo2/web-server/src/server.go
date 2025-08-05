package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var (
	err    error
	config Config
	db     *sqlx.DB
)

func getExitSignal() chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	return sig
}

func startWebServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/get/albums", getAllAlbums)
	mux.HandleFunc("/get/album", getAlbum)
	mux.HandleFunc("/add/album", addAlbumHandler)
	mux.HandleFunc("/add/review", addReviewHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
	log.Info("Listening on port 8080")
}

func main() {
	log.SetLevel(log.DebugLevel)

	config = getConfig()
	dsn := getDSN()

	log.Info("Connecting to mysql database...")
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Error("Error while connecting to database: ", err.Error())
		os.Exit(1)
	}
	defer db.Close()
	log.Info("Successfully connected to the mysql database.")

	exitSig := getExitSignal()

	go startWebServer()

	<-exitSig
}
