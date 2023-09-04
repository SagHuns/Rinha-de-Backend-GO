package models

import(
	"github.com/SagHuns/Rinha-de-Backend-GO/db"
)

func ContadorPessoas() (contagem int, err error) {
	// Primeiro passo é tentar abrir uma conexão com o banco de dados
	conn := db.GetDB()
	
	// Fecha o DB quando a operação encerrar
	defer conn.Close()

	row := conn.QueryRow(`SELECT COUNT(*) FROM pessoas`)
	err = row.Scan(&contagem)
	return

}