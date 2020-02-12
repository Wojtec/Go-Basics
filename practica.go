package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))

}

func mainHandler(w http.ResponseWriter, r *http.Request) {
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
