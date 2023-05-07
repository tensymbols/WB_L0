package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"os"
)

func main() {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	//	fmt.Println(err)
	defer db.Close()
	var str string
	err = db.QueryRow("select `hello`").Scan(&str)

	fmt.Println(err, str)
}
