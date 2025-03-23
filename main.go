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

	//USER RELATED ENDPOINTS
	//Create User
	mux.HandleFunc("POST /v1/users/", createUserHandler)
	//Get User
	mux.HandleFunc("GET /v1/users/", getUserHandler)
	//Delete User
	mux.HandleFunc("DELTE /v1/users/", deleteUserHandelr)
	//Update User
	mux.HandleFunc("PUT /v1/users/", updateUserHandler)

	//POST RELATED ENDPOINTS
	//Create Post
	mux.HandleFunc("POST /v1/posts/", createPostHandler)
	//Get Post
	mux.HandleFunc("GET /v1/posts/", getPostHandler)
	//Delete Post
	mux.HandleFunc("DELETE /v1/posts/", deletePostHandler)
	//Update Post
	mux.HandleFunc("PUT /v1/posts/", updatePostHandler)

	srv := http.Server{
		Addr:    "localhost:" + port,
		Handler: mux,
	}
	fmt.Println("Listening and Serving on port: ", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
