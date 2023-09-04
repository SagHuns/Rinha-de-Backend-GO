package main

import (
    "log"
    "net/http"
    
    "github.com/SagHuns/Rinha-de-Backend-GO/handlers"
    "github.com/SagHuns/Rinha-de-Backend-GO/db"
)

func main() {
    db.InitDB()
    db.InitSchema()
    PORT := ":8080"
    log.Print("Server initialized at port ", PORT)
    http.HandleFunc("/pessoas", handlers.Pessoas)
    http.HandleFunc("/pessoas/", handlers.Get)
    http.HandleFunc("/contagem-pessoas", handlers.Count)
    log.Fatal(http.ListenAndServe(PORT, nil))
}
