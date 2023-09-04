package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SagHuns/Rinha-de-Backend-GO/models"
)

func Post(w http.ResponseWriter, r *http.Request) {
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
