package models

import (
	"context"
	"strconv"

	"github.com/SagHuns/Rinha-de-Backend-GO/db"
)

func ContadorPessoas() (contagem int, err error) {
	rdb := db.GetRedis()
	val, err := rdb.Get(context.Background(), "contador_pessoas").Result()
	// Se não houver erro, significa que o valor foi encontrado no Redis
	if err == nil {
		// Se houver erro, significa que o valor não foi encontrado no Redis
		contagem, err = strconv.Atoi(val)
		if err == nil {
			return
		}
	}

	pool := db.GetDB()
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return 0, err
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(), `SELECT COUNT(*) FROM pessoas`)
	err = row.Scan(&contagem)
	// Se não houver erro, significa que o valor foi encontrado no banco de dados
	if err == nil {
		// Setando o valor no Redis para que não seja necessário consultar o banco de dados novamente
		err = rdb.Set(context.Background(), "contador_pessoas", contagem, 0).Err()
	}
	return
}
