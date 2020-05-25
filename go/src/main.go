package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":6800", nil)

	
}

func handler(w http.ResponseWriter, r *http.Request) {
	greetingMessage := greeting("Code.education Rocks!")
    fmt.Fprintf(w, greetingMessage)
}

func greeting(message string) string {
	htmlBoldMessage := "<b>" + message + "</b>"

	return htmlBoldMessage;
}
