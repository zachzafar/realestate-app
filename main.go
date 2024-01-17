package main

import (
	"flag"
	"os"

	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"openlettings.com/db"
	"openlettings.com/server"
)

var logger = logrus.New()

func init() {
	logger.SetLevel(logrus.InfoLevel)

	logger.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.Create("logfile.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger.SetOutput(file)
}

func main() {
	db, err := db.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	listenAddress := flag.String("listenaddr", ":3000", "the server address")
	flag.Parse()
	server := server.NewServer(*listenAddress, db, logger)
	log.Fatal(server.Start())

}
