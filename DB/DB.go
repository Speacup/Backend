package DB

import (
	"database/sql"

	_ "github.com/lib/pq"

	"fmt"
)

func DBConnect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DBhost, DBport, DBuser, DBpassword, DBname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Database connection successful")
	}
	return db

}
func DBPing(D *sql.DB) {
	err := D.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Print("Database Ping Successful")
	}
}
