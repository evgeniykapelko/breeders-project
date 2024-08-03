package main

import (
	"breeders/models"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
	Models      models.Models
}

type appConfig struct {
	useCache bool
	dsn      string
}

func main() {
	app := &application{
		templateMap: make(map[string]*template.Template),
	}

	flag.BoolVar(&app.config.useCache, "cache", false, "use cache for templates")
	flag.StringVar(&app.config.dsn, "dsn", "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "DSN")
	flag.Parse()

	db, err := initMySQLDB(app.config.dsn)
	if err != nil {
		log.Panic(err)
	}

	app.Models = *models.New(db)

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "Hello, World!")
	//})

	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       time.Second * 30,
		ReadTimeout:       time.Second * 30,
		ReadHeaderTimeout: time.Second * 30,
		WriteTimeout:      time.Second * 30,
	}

	fmt.Println("Starting web application on port", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	//err := http.ListenAndServe(port, nil)
	//if err != nil {
	//	log.Panic(err)
	//}
}
