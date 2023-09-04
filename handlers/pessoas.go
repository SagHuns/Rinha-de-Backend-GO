package handlers

import "net/http"

func Pessoas(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        Search(w, r)
    } else if r.Method == http.MethodPost {
        Post(w, r)
    }
}
