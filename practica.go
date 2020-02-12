package main

import (
	"encoding/xml"
	"log"
	"net/http"
)

//type struct

type Profile struct {
	Name    string
	Hobbies []string
}

func main() {
	http.HandleFunc("/", htmlHandler)
	http.HandleFunc("/xml", xmlHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))

}

func xmlHandler(w http.ResponseWriter, r *http.Request) {

	profile := Profile{"Wojtek", []string{"Travels", "Programming"}}

	x, err := xml.MarshalIndent(profile, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "text/xml; charset=utf8")
	w.WriteHeader(http.StatusOK)
	w.Write(x)
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`<html>
		<head><title>Simple server</title></head>
		<body>
		<h1>Hello </h1>
		<p>First start website in go</p>
		<body>
		</html>
		`))

}
