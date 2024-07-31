package main

import (
	"log"
	"net/http"

	"github.com/wendao2000/simopi/app"
)

func main() {
	http.HandleFunc("/config", app.GetConfig)
	http.HandleFunc("/create", app.CreateConfig)
	http.HandleFunc("/delete", app.DeleteConfig)
	http.HandleFunc("/", app.MatchmakeConfig)

	log.Print("Listening to :4444")
	http.ListenAndServe(":4444", nil)
}
