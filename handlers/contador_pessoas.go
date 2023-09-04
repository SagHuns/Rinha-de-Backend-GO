package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SagHuns/Rinha-de-Backend-GO/models"
)

func Count(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        contagem, err := models.ContadorPessoas()
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(contagem)

    } else {
        w.WriteHeader(http.StatusNotFound)
    }
}
