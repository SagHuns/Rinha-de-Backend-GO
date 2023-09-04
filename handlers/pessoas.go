package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
	"time"
    "log"

	"github.com/SagHuns/Rinha-de-Backend-GO/models"
	"github.com/google/uuid"
)

func PessoasHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        PessoasSearchHandler(w, r)
    } else if r.Method == http.MethodPost {
        PessoasPostHandler(w, r)
    }
}

func PessoasPostHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")

    var PessoaJson models.Pessoa
    const nascimento_valido = "2006-01-02"

    var error1 error = json.NewDecoder(r.Body).Decode(&PessoaJson)

    if error1 != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    if PessoaJson.Apelido == "" {
        w.WriteHeader(http.StatusUnprocessableEntity)
        w.Write([]byte("Digite um Apelido não nulo!"))
        return
    }

    if PessoaJson.Nome == "" {
        w.WriteHeader(http.StatusUnprocessableEntity)
        w.Write([]byte("Digite um Nome não nulo!"))
        return
    }

    if PessoaJson.Nascimento == "" {
        w.WriteHeader(http.StatusUnprocessableEntity)
        w.Write([]byte("Digite um Nascimento não nulo!"))
        return
    }

    _, erro := time.Parse(nascimento_valido, PessoaJson.Nascimento)
    if erro != nil {
        w.WriteHeader(http.StatusUnprocessableEntity)
        w.Write([]byte("Digite a data de nascimento no formato AAAA-MM-DD"))
        return
    }

    id, err := models.Create(PessoaJson)
    if err != nil {
        if err.Error() == "apelido já existe" {
            w.WriteHeader(http.StatusUnprocessableEntity)
            w.Write([]byte(err.Error()))
            return
        } else {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte("Erro ao criar pessoa"))
            return
        }
    }

    url := "/pessoas/" + id.String()
    PessoaJson.Id = id

    w.Header().Set("Location", url)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(PessoaJson)

}


func PessoasGetHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        idStr := strings.TrimPrefix(r.URL.Path, "/pessoas/")
        id, err := uuid.Parse(idStr)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("Invalid UUID"))
            return
        }

        pessoa, err := models.Get(id)
        log.Println(err)
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

func PessoasSearchHandler(w http.ResponseWriter, r *http.Request) {
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


func PessoasContagemHandler(w http.ResponseWriter, r *http.Request) {
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
