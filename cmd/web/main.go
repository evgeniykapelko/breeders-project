package main

import (
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
}

type appConfig struct {
	useCache bool
}

func main() {
	app := &application{
		templateMap: make(map[string]*template.Template),
	}

	flag.BoolVar(&app.config.useCache, "cache", false, "use cache for templates")
	flag.Parse()

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

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	//err := http.ListenAndServe(port, nil)
	//if err != nil {
	//	log.Panic(err)
	//}
}
