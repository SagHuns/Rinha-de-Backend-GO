package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
    "github.com/google/uuid"
    "strings"

)

type Pessoa struct {
    Apelido string
    Nome string
    Nascimento string
    Stack []string
    Id uuid.UUID
}

var Database []Pessoa  // Cria DB simples de pessoas

func main() {
    PORT := ":8080"
    log.Print("Server initialized at port ", PORT)
    http.HandleFunc("/pessoas", pessoasHandler)
    http.HandleFunc("/pessoas/", pessoasGetHandler)
    log.Fatal(http.ListenAndServe(PORT, nil))  // Ouve os endpoints
}


// Um meio para saber se o /pessoas é para fazer uma busca e dispôr atributos
// ou se é para fazer um POST
func pessoasHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        pessoasSearchHandler(w, r)

    } else if r.Method == http.MethodPost {
        pessoasPostHandler(w, r)
    }
}
func pessoasPostHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")  // Mudando o formato principal de envio para JSON

    var PessoaJson Pessoa
    const nascimento_valido = "2006-01-02"

    var error1 error = json.NewDecoder(r.Body).Decode(&PessoaJson)  // Caso ocorra um erro na hora de decodificar o JSON

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
    // Cria url com esse UUID
    url := "/pessoas/" + id.String()
    PessoaJson.Id = id
    Database = append(Database, PessoaJson)

    // Write the created URL to the response header
    w.Header().Set("Location", url)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(PessoaJson)

}

func pessoasGetHandler(w http.ResponseWriter, r *http.Request) {
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

func pessoasSearchHandler(w http.ResponseWriter, r *http.Request) {
    termo := r.URL.Query().Get("t")
    if termo == "" {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Termo de busca não informado"))
        return
    }

    var resultados []Pessoa
    for _, pessoa := range Database {
        if strings.Contains(pessoa.Apelido, termo) || strings.Contains(pessoa.Nome, termo) {
            resultados = append(resultados, pessoa)
            continue
        }
        for _, stack := range pessoa.Stack {
            if strings.Contains(stack, termo) {
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

