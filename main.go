package main

import (
    "log"
    "net/http"
    "github.com/SagHuns/Rinha-de-Backend-GO/handlers"
)

func main() {
    PORT := ":8080"
    log.Print("Server initialized at port ", PORT)
    http.HandleFunc("/pessoas", handlers.PessoasHandler)
    http.HandleFunc("/pessoas/", handlers.PessoasGetHandler)
    log.Fatal(http.ListenAndServe(PORT, nil))
}
