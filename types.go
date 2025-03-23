package main

type ErrResponsePayload struct {
	Message string `json:"error_message"`
	Extras  string `json:"additonal_info"`
}
