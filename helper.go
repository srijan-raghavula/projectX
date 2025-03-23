package main

import (
	"encoding/json"
	"fmt"
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

func respondWithErr(w http.ResponseWriter, code int, err error, msg string) {
	fmt.Println("Responding with error:", err)

	payload, err := json.Marshal(ErrResponsePayload{
		Message: err.Error(),
		Extras:  msg,
	})
	if err != nil {
		w.Header().Add("Content-Type", "text/html")
		respondWithErr(w, http.StatusInternalServerError, err, "failed to marshal error message")
		return
	}

	w.WriteHeader(code)
	w.Write(payload)
}

func respondWithPayload(w http.ResponseWriter, code int, payload any) {
	fmt.Println("Responding with payload")

	writeData, err := json.Marshal(payload)
	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, err, "Failed to marshal the payload")
		return
	}

	w.WriteHeader(code)
	w.Write(writeData)
}
