package models


import(
	"github.com/SagHuns/Rinha-de-Backend-GO/db"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func Get(id uuid.UUID) (pessoa Pessoa, err error) {
	conn := db.GetDB()
	
	row := conn.QueryRow(`SELECT * FROM pessoas WHERE id=$1`, id)
	// Variáveis temporárias para serem copiadas para as variáveis da struct pessoa
	var stack []string
	var tempId uuid.UUID
	
	err = row.Scan(&tempId, &pessoa.Apelido, &pessoa.Nome, &pessoa.Nascimento, pq.Array(&stack))
	pessoa.Stack = stack
	pessoa.Id = tempId

	return
}
