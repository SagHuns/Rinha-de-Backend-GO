package models

import (
	"errors"
	"log"

	"github.com/SagHuns/Rinha-de-Backend-GO/db"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func Create(pessoa Pessoa) (id uuid.UUID, err error) {
	id = uuid.New()
	
	conn := db.GetDB()
 
	// Checar se apelido ja existe
	var count int
	err = conn.QueryRow("SELECT COUNT(*) FROM pessoas WHERE apelido = $1", pessoa.Apelido).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count > 0 {
		err = errors.New("apelido já existe")
		return
	}

	sql := `INSERT INTO pessoas (id, apelido, nome, nascimento, stack) VALUES ($1, $2, $3, $4, $5)`
	// pq.Array() é para converter no formato que o postgres armazena os arrays.
	_, err = conn.Exec(sql, id.String(), pessoa.Apelido, pessoa.Nome, pessoa.Nascimento, pq.Array(pessoa.Stack))
	if err != nil {
		log.Fatal(err)
	}
	return
}
