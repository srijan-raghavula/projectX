package main

import (
	"net/http"
)

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	tempHandler(w, r)
}

func getPostHandler(w http.ResponseWriter, r *http.Request) {
	tempHandler(w, r)
}

func deletePostHandler(w http.ResponseWriter, r *http.Request) {
	tempHandler(w, r)
}

func updatePostHandler(w http.ResponseWriter, r *http.Request) {
	tempHandler(w, r)
}
