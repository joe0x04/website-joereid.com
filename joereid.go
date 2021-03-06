package main

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

var AID string

type IndexData struct {
	Title      string
	Adsense_ID string // via env variable
}

/*
 * Handle requests for the index page
 */
func index(w http.ResponseWriter, r *http.Request) {
	data := IndexData{
		Title:      "Joe Reid",
		Adsense_ID: AID,
	}

	t, _ := template.ParseFiles("template/index.tmpl")
	t.Execute(w, data)

	log.Println(r.Method, r.RemoteAddr, r.URL.Path, r.UserAgent())
}

func main() {
	r := mux.NewRouter()

	// catch any requests for "real" files (css, js, images, etc)
	r.PathPrefix("/assets/").Handler(http.FileServer(http.Dir("static/")))

	// catch index request
	r.HandleFunc("/", index)

	http.ListenAndServe(":8087", r)
}

func init() {
	logfile, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.SetOutput(logfile)

	AID = os.Getenv("AID")
}
