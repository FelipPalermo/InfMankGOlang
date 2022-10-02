package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {

	result, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

	return result
}

func RP(palavra, horario string) string {

	insert := fmt.Sprintf("INSERT INTO dict(palavra, horario) VALUES ('%s', '%s')", palavra, horario)

	return insert
}

func RN(maximo int) int {
	number := rand.Intn(maximo)
	return number
}

func main() {

	start := time.Now()

	db, err := sql.Open("mysql", "root:root@/InfMank")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	rand.Seed(time.Now().UnixNano())

	var letra [1]string
	var holder string
	var palavra string

	alfabeto := [27]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", " ", "z"}

	for {

		palavra = ""
		holder = ""

		for i := 0; i < RN(10)+2; i++ {

			letra[0] = alfabeto[RN(27)]
			holder = holder + letra[0]

		}

		palavra = holder
		fmt.Println(palavra)

		time.Sleep(time.Nanosecond)

		hora_palavra := fmt.Sprint(time.Now())
		exec(db, RP(palavra, hora_palavra))

	}

	duration := time.Since(start)
	fmt.Print(duration)

}
