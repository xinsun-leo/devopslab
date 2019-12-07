package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func say(word string) string {
	title := "Second Hello world, this is"
	return title + " " + word
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Say something!")
	fmt.Fprintf(w, "%s\n", say("BLUE-IBM!!!"))
}

func main() {
	log.Print("Say something started.")

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
