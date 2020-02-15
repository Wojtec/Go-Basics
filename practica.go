package main

import (
	"encoding/json"
	"encoding/xml"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

//type struct

type Profile struct {
	Name    string
	Hobbies []string
}

func main() {
	http.HandleFunc("/", htmlHandler)
	http.HandleFunc("/xml", xmlHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/plain", plainHandler)
	http.HandleFunc("/image", imageHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))

}
func imageHandler(w http.ResponseWriter, r *http.Request) {
	img, err := os.Open("scooby.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer img.Close()
	w.Header().Set("Content-type", "image/jpeg; charset=utf8")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, img)
}

func plainHandler(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Wojtek", []string{"Travels", "Programming"}}

	plain, err := json.MarshalIndent(profile, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "text/plain; charset=utf8")
	w.WriteHeader(http.StatusOK)
	w.Write(plain)

}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusOK)
	profile := Profile{"Wojtek", []string{"Travels", "Programming"}}

	js, err := json.MarshalIndent(profile, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}
func xmlHandler(w http.ResponseWriter, r *http.Request) {

	profile := Profile{"Wojtek", []string{"Travels", "Programming"}}

	x, err := xml.MarshalIndent(profile, "", "")
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
	t, e := template.ParseFiles("template.html")
	t.Execute(w, e)
}
