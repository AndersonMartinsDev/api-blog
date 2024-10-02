package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //Driver
)

func Conectar() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, erro := sql.Open("postgres", psqlconn)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db, nil
}

func GetConnectionDatabase() *sql.DB {
	banco, erro := Conectar()

	if erro != nil {
		panic("Conexão Não estabelecida com o Banco de dados!")
	}
	return banco
}
