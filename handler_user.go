package main

import "net/http"

func handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
