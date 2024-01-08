package main

import (
	"flag"

	"log"

	_ "github.com/mattn/go-sqlite3"
	"openlettings.com/db"
	"openlettings.com/server"
)

func main() {
	db, err := db.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	listenAddress := flag.String("listenaddr", ":3000", "the server address")
	flag.Parse()
	server := server.NewServer(*listenAddress, db)
	log.Fatal(server.Start())

}
