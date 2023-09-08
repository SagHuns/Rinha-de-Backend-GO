package models

import (
	"context"
	"encoding/json"
	"log"

	"github.com/SagHuns/Rinha-de-Backend-GO/db"
	"github.com/jackc/pgtype"
)

func SearchPessoas(termo string) ([]Pessoa, error) {
	rdb := db.GetRedis()
	val, err := rdb.Get(context.Background(), "search_pessoas:"+termo).Result()
	if err == nil {
		var resultados []Pessoa
		err = json.Unmarshal([]byte(val), &resultados)
		if err == nil {
			return resultados, nil
		}
	}

	pool := db.GetDB()
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), `SELECT id, apelido, nome, nascimento, stack FROM pessoas 
	WHERE LOWER(apelido) LIKE $1 OR LOWER(nome) LIKE $1 OR $1 = ANY(stack)`, "%"+termo+"%")
	if err != nil {
		return nil, err
	}

	var resultados []Pessoa
	for rows.Next() {
		var pessoa Pessoa
		var stack pgtype.TextArray
		err = rows.Scan(&pessoa.Id, &pessoa.Apelido, &pessoa.Nome, &pessoa.Nascimento, &stack)
		if err != nil {
			continue
		}
		
		pessoa.Stack = make([]string, len(stack.Elements))
		for i, element := range stack.Elements {
			pessoa.Stack[i] = element.String
		}

		resultados = append(resultados, pessoa)
	}

	data, err := json.Marshal(resultados)
	if err == nil {
		err = rdb.Set(context.Background(), "search_pessoas:"+termo, data, 0).Err()
	}

	return resultados, nil
}
