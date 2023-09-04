package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/SagHuns/Rinha-de-Backend-GO/models"
	"github.com/google/uuid"
)

func Get(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        idStr := strings.TrimPrefix(r.URL.Path, "/pessoas/")
        id, err := uuid.Parse(idStr)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("Invalid UUID"))
            return
        }

        pessoa, err := models.Get(id)
        if err == sql.ErrNoRows {
            w.WriteHeader(http.StatusNotFound)
            w.Write([]byte("Pessoa não encontrada"))
            return
        } else if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(pessoa)

    } else {
        w.WriteHeader(http.StatusNotFound)
    }
}
