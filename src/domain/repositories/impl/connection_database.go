package impl

import (
	"apiblog/src/infrastructure/database"
	"database/sql"
)

func GetConnectionDatabase() *sql.DB {
	banco, erro := database.Conectar()

	if erro != nil {
		panic("Conexão Não estabelecida com o Banco de dados!")
	}
	return banco
}
