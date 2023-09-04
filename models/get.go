package models


import(
	"github.com/SagHuns/Rinha-de-Backend-GO/db"
	"github.com/google/uuid"
)

func Get(id uuid.UUID) (pessoa Pessoa, err error) {
	// Primeiro passo é tentar abrir uma conexão com o banco de dados
	conn := db.GetDB()
 
	// Fecha o DB quando a operação encerrar
	defer conn.Close()
	
	row := conn.QueryRow(`SELECT * FROM pessoas WHERE id=$1`, id)

	err = row.Scan(&pessoa.Apelido, &pessoa.Nome, &pessoa.Nascimento, &pessoa.Stack)

	return
}