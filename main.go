package main

import (
	"flag"
	"os"

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

	if err != nil && os.Getenv("APP_ENV") != "PRODUCTION" {
		panic(err)
	}

	var logger = logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.SetOutput(os.Stdout)

	err = runMigrations()

	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("successfully updated db")

	db, err := db.InitDB()

	if err != nil {
		logger.Fatal(err)
	}

	InfoStore, err := db.InitStore()

	if err != nil {
		logger.Fatal(err)
	}

	listenAddress := flag.String("listenaddr", ":"+os.Getenv("PORT"), "the server address")
	flag.Parse()
	server := server.NewServer(*listenAddress, db, logger, *InfoStore)
	logger.Fatal(server.Start())

}

func runMigrations() error {
	migrationsPath := "file://db/migrations"
	dbsource := os.Getenv("DB_URL")

	migration, err := migrate.New(migrationsPath, dbsource)

	if err != nil {
		return err
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
