package handlers


import (
    "encoding/json"
    "net/http"
    "time"
    "github.com/google/uuid"
    "strings"
    "github.com/SagHuns/Rinha-de-Backend-GO/models"
)

var Database []models.Pessoa

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

    for _, pessoa := range Database {
        if pessoa.Apelido == PessoaJson.Apelido {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("Apelido já existe!"))
            return
        }
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

    id := uuid.New()
    url := "/pessoas/" + id.String()
    PessoaJson.Id = id
    Database = append(Database, PessoaJson)

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

        for _, pessoa := range Database {
            if pessoa.Id == id {
                json.NewEncoder(w).Encode(pessoa)
                return
            }
        }

        w.Write([]byte("Pessoa not found"))
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

    var resultados []Pessoa
    for _, pessoa := range Database {
        if strings.Contains(strings.ToLower(pessoa.Apelido), termo) || strings.Contains(strings.ToLower(pessoa.Nome), termo) {
            resultados = append(resultados, pessoa)
            continue
        }
        for _, stack := range pessoa.Stack {
            if strings.Contains(strings.ToLower(stack), termo) {
                resultados = append(resultados, pessoa)
                break
            }
        }
    }

    if len(resultados) > 50 {
        resultados = resultados[:50]
    }

    json.NewEncoder(w).Encode(resultados)
}