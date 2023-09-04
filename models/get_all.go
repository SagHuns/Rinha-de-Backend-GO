package models


import(
	"github.com/SagHuns/Rinha-de-Backend-GO/db"
)

func GetAll() (pessoas[] Pessoa, err error) {
	// Primeiro passo é tentar abrir uma conexão com o banco de dados
	conn := db.GetDB()
 
	// Fecha o DB quando a operação encerrar
	defer conn.Close()
	
	rows, err := conn.Query(`SELECT * FROM pessoas`)
	if err != nil {
		return
	}

	for rows.Next(){
		var pessoa Pessoa
		err = rows.Scan(&pessoa.Apelido, &pessoa.Nome, &pessoa.Nascimento, &pessoa.Stack)
		if err != nil {
			continue
		}

		pessoas = append(pessoas, pessoa)
	}

	return
}