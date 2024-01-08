package db

import (
	"database/sql"

	"github.com/michaeljs1990/sqlitestore"
)

const filename = "sqlite.db"

const laterUse = `CREATE TABLE IF NOT EXISTS properties (
  property_id INTEGER PRIMARY KEY AUTOINCREMENT,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  property_type VARCHAR(50) NOT NULL,
  address VARCHAR(255) NOT NULL,
  city VARCHAR(100) NOT NULL,
  state VARCHAR(50) NOT NULL,
  zip_code VARCHAR(20) NOT NULL,
  neighborhood VARCHAR(255),
  size INTEGER,
  bedrooms INTEGER,
  bathrooms INTEGER,
  year_built INTEGER,
  flooring_type VARCHAR(50),
  appliances TEXT[],
  price DECIMAL(10, 2),
  currency VARCHAR(3),
  payment_terms VARCHAR(255),
  amenities TEXT[],
  features TEXT[],
  availability BOOLEAN DEFAULT true,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  user_id INT REFERENCES users(user_id) ON DELETE CASCADE
);`

const initQuery = `
		CREATE TABLE IF NOT EXISTS users (
  user_id INTEGER PRIMARY KEY AUTOINCREMENT,
  username VARCHAR(255) NOT NULL UNIQUE,
  password_hash VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  role VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS properties (
  property_id INTEGER PRIMARY KEY AUTOINCREMENT,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  property_type VARCHAR(50) NOT NULL,
  address VARCHAR(255) NOT NULL,
  city VARCHAR(100) NOT NULL,
  size INTEGER,
  bedrooms INTEGER,
  bathrooms INTEGER,
  year_built INTEGER,
  price DECIMAL(10, 2),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  user_id INT REFERENCES users(user_id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS media (
  media_id SERIAL PRIMARY KEY,
  property_id INT REFERENCES properties(property_id) ON DELETE CASCADE,
  url VARCHAR(255) NOT NULL
);
`

type Database struct {
	db           *sql.DB
	sessionStore *sqlitestore.SqliteStore
}

func InitDB() (*Database, error) {
	db, err := sql.Open("sqlite3", filename)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(initQuery)

	if err != nil {
		return nil, err
	}

	store, err := sqlitestore.NewSqliteStore("sqlite.db", "sessions", "/", 3600, []byte("thisisabadsecret"))
	if err != nil {
		return nil, err
	}

	return &Database{db: db, sessionStore: store}, nil

}
