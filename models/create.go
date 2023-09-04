package models

import(
	"github.com/SagHuns/Rinha-de-Backend-GO/db"
	"github.com/google/uuid"
)

func Create(pessoa Pessoa) (id uuid.UUID, err error) {
	id = uuid.New()
	
	// Primeiro passo é tentar abrir uma conexão com o banco de dados
	conn := db.GetDB()
 
	// Fecha o DB quando a operação encerrar
	defer conn.Close()
	
	sql := `INSERT INTO pessoas (id, apelido, nome, nascimento, stack) VALUES ($1, $2, $3, $4, $5)`

	_, err = conn.Exec(sql, id.String(), pessoa.Apelido, pessoa.Nome, pessoa.Nascimento, pessoa.Stack)
	return
}
