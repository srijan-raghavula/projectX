package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "8888"

func main() {
	fmt.Println("Starting the main server")

	staticFS := http.FileServer(http.Dir("./src/static"))
	mux := http.NewServeMux()

	mux.Handle("/v1/static/", http.StripPrefix("/v1/static", staticFS))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		staticFS.ServeHTTP(w, r)
	})

	mux.HandleFunc("/v1/test/", tempHandler)

	srv := http.Server{
		Addr:    "localhost:" + port,
		Handler: mux,
	}
	fmt.Println("Listening and Serving on port: ", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
