package main

import (
	"breeders/models"
	"log"
	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
	testApp = application{}
	dsn := "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s"

	db, err := initMySQLDB(dsn)

	if err != nil {
		log.Fatal(err)
	}

	testApp = application{
		DB:     db,
		Models: *models.New(db),
	}

	os.Exit(m.Run())
}
