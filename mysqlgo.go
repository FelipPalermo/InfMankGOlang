package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Funcao que recebe uma referencia de database de sql e uma string
// e retorna o resultado da query
func exec(db *sql.DB, texto string) sql.Result {

	// db.Exc retorna 2 variaveis, resultado e erro
	result, err := db.Exec(texto)

	// se erro for diferente de nada
	// entao a gente trata o erro
	if err != nil {
		panic(err)
	}

	return result
}

// funcao para que recebe palavra e horario que serao adicionados
func RP(palavra, horario string) string {

	// Sprintf retorna uma string formatada
	insert := fmt.Sprintf("INSERT INTO dict(palavra, horario) VALUES ('%s', '%s')", palavra, horario)

	return insert
}

func main() {

	// inicia uma instancia de sql "mysql", "nome de usuario:senha@/database"
	// se nao quisermos nenhuma database previa usamos @/ sem nada
	db, err := sql.Open("mysql", "root:root@/")

	// Tratamento de erro
	if err != nil {
		panic(err)
	}

	// Fechar mysql quando a funcao main for fechada
	defer db.Close()

	exec(db, "use InfMank")
	exec(db, RP("felipe", "1 da manha"))

}
