package models

import (
	"context"
	"encoding/json"
	"log"

	"github.com/SagHuns/Rinha-de-Backend-GO/db"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

func Get(id uuid.UUID) (pessoa Pessoa, err error) {
	rdb := db.GetRedis()
	val, err := rdb.Get(context.Background(), "pessoa:"+id.String()).Result()
	if err == nil {
		err = json.Unmarshal([]byte(val), &pessoa)
		if err == nil {
			return
		}
	}

	pool := db.GetDB()
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(), `SELECT * FROM pessoas WHERE id=$1`, id)
	var stack pgtype.TextArray
	var tempId uuid.UUID

	err = row.Scan(&tempId, &pessoa.Apelido, &pessoa.Nome, &pessoa.Nascimento, &stack)
	pessoa.Stack = make([]string, len(stack.Elements))
	for i, element := range stack.Elements {
		pessoa.Stack[i] = element.String
	}
	pessoa.Id = tempId
	
	if err == nil {
		data, err := json.Marshal(pessoa)
		if err == nil {
			err = rdb.Set(context.Background(), "pessoa:"+id.String(), data, 0).Err()
		}
	}

	return
}
