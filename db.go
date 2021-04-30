package main

import (
	"database/sql"
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	tb "gopkg.in/tucnak/telebot.v2"
)

var dbobj *sql.DB

func init() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}
	log.Printf("Connecting to uri %s\n", getConnectionURI())
	dbobj, err = sql.Open("postgres", getConnectionURI())
	if err != nil {
		log.Fatalln(err)
	}
	err = dbobj.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

func getConnectionURI() string {
	return os.Getenv("NASA_APOD_TELEGRAM_BOT_DATABASE_URI")
}

func getData() []tb.User {
	tx, err := dbobj.Begin()
	if err != nil {
		log.Fatalln(err)
	}
	defer tx.Rollback()
	rows, err := dbobj.Query("select * from tg_user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	users := []tb.User{}
	for rows.Next() {
		user := tb.User{}
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username,
			&user.LanguageCode, &user.IsBot, &user.CanJoinGroups, &user.CanReadMessages,
			&user.SupportsInline)
		if err != nil {
			log.Fatalln(err)
			defer tx.Rollback()
			return []tb.User{}
		}
		users = append(users, user)
	}
	tx.Commit()
	return users
}

func addUser(user *tb.User) {
	tx, err := dbobj.Begin()
	if err != nil {
		log.Fatalln(err)
	}
	defer tx.Rollback()
	stmt, err := dbobj.Prepare("INSERT INTO tg_user VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ON CONFLICT DO NOTHING")
	if err != nil {
		log.Fatalln(err)
		return
	}

	_, err = stmt.Exec(&user.ID, &user.FirstName, &user.LastName, &user.Username,
		&user.LanguageCode, &user.IsBot, &user.CanJoinGroups, &user.CanReadMessages,
		&user.SupportsInline)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Printf("Inserted new user %s\n", user.Username)
	tx.Commit()
}
