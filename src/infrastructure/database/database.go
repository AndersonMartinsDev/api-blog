package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

func InitalStrucuture() {
	banco, erro := Conectar()
	if erro != nil {
		panic("Problema ao criar tabelas")
	}

	banco.Exec("CREATE DATABASE apiblog")

	config := &postgres.Config{SchemaName: "public"}
	driver, err := postgres.WithInstance(banco, config)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(config.DatabaseName)

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		config.DatabaseName,
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	defer banco.Close()
}
