package main

import (
	"net/http"
)

func tempHandler(w http.ResponseWriter, _ *http.Request) {
	// This is a function to use a placeholder handler so that the LSP don't scream at me
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width initial-scale=1.0">
            <title>Placeholder</title>
        </head>
        <body>
            <h1>Placeholder lol</h1>
        </body>
        </html>
        `))
}
