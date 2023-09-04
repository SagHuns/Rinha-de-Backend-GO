package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/SagHuns/Rinha-de-Backend-GO/models"
)

func Search(w http.ResponseWriter, r *http.Request) {
    termo := r.URL.Query().Get("t")
    if termo == "" {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Termo de busca não informado"))
        return
    }

    termo = strings.ToLower(termo)

    resultados, err := models.SearchPessoas(termo)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Erro ao buscar pessoas"))
        return
    }

    if len(resultados) > 50 {
        resultados = resultados[:50]
    }

    json.NewEncoder(w).Encode(resultados)
}
