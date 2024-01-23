package main

import (
	"flag"
	"fmt"
	"os"

	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"openlettings.com/db"
	"openlettings.com/server"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
	var logger = logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	logger.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.Create("logfile.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger.SetOutput(file)

	err = runMigrations()

	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("successfully updated db")

	db, err := db.InitDB()

	if err != nil {
		logger.Fatal(err)
	}

	listenAddress := flag.String("listenaddr", ":3000", "the server address")
	flag.Parse()
	server := server.NewServer(*listenAddress, db, logger)
	logger.Fatal(server.Start())

}

func runMigrations() error {
	migrationsPath := "file://db/migrations"
	dbsource := os.Getenv("DB_URL")
	fmt.Println(dbsource)
	migration, err := migrate.New(migrationsPath, dbsource)

	if err != nil {
		return err
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
