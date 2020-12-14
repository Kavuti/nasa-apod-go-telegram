package main

import (
	"database/sql"
	"os"
	// _ "github.com/lib/pq"
)

var dbobj sql.DB

type User struct {
	id       int
	username string
}

func init() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Println("No .env file found")
	// }
	// dbobj, err := sql.Open("postgres", getConnectionURI())
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// err = dbobj.Ping()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}

func getConnectionURI() string {
	return os.Getenv("NASA_APOD_TELEGRAM_BOT_DATABASE_URI")
}

func getData() {
	// rows, err := dbobj.Query("select * from users")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	err := rows.Scan()
	// }
}
