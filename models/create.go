package models

import (
	"context"
	"errors"
	"log"

	"github.com/SagHuns/Rinha-de-Backend-GO/db"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

func Create(pessoa Pessoa) (id uuid.UUID, err error) {
	id = uuid.New()
	
	pool := db.GetDB()  // Criando um pgxpool para explorar o paralelismo no banco de dados
	conn, err := pool.Acquire(context.Background())  // Criando uma conexão com o banco de dados
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Release()  // Fechando a conexão com o banco de dados quando a função terminar
 
	// Checar se apelido ja existe
	// Como eu implementei os indexes antes, não preciso alterar esse trecho de código
	var count int
	err = conn.QueryRow(context.Background(),"SELECT COUNT(*) FROM pessoas WHERE apelido = $1", pessoa.Apelido).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count > 0 {
		err = errors.New("apelido já existe")
		return
	}

	sql := `INSERT INTO pessoas (id, apelido, nome, nascimento, stack) VALUES ($1, $2, $3, $4, $5)`
	stack := &pgtype.TextArray{}
	stack.Set(pessoa.Stack)
	
	_, err = conn.Exec(context.Background(), sql, id.String(), pessoa.Apelido, pessoa.Nome, pessoa.Nascimento, stack)
	if err != nil {
		log.Fatal(err)
	}
	return
}
