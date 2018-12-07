package models

import (
	"log"
	"net/url"
)

var TestDB *DB

func init() {
	userPassword := url.UserPassword("postgres", "postgres")
	URI := new(url.URL)
	URI.Scheme = "postgres"
	URI.User = userPassword
	URI.Host = "localhost"
	URI.Path = "apdf_db"
	URI.RawQuery = "sslmode=disable"
	db, err := NewDB(URI.String())
	if err != nil {
		log.Fatalf("\n DB initialization failed: %v",
			err)
	}
	TestDB = db
}
