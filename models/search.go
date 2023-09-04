package models

import (
	"github.com/SagHuns/Rinha-de-Backend-GO/db"
	"github.com/lib/pq"
)

func SearchPessoas(termo string) ([]Pessoa, error) {
	conn := db.GetDB()
	rows, err := conn.Query(`SELECT id, apelido, nome, nascimento, stack FROM pessoas 
	WHERE LOWER(apelido) LIKE $1 OR LOWER(nome) LIKE $1 OR $1 = ANY(stack)`, "%"+termo+"%")
	
	if err != nil {
		return nil, err
	}

	var resultados []Pessoa
	for rows.Next() {
		var pessoa Pessoa
		err = rows.Scan(&pessoa.Id, &pessoa.Apelido, &pessoa.Nome, &pessoa.Nascimento, pq.Array(&pessoa.Stack))
		if err != nil {
			continue
		}

		resultados = append(resultados, pessoa)
	}

	return resultados, nil
}
