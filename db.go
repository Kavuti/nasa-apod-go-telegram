package main

import (
	"database/sql"
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gopkg.in/tucnak/telebot.v2"
)

var dbobj sql.DB

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}
	dbobj, err := sql.Open("postgres", getConnectionURI())
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

func getData() []User {
	rows, err := dbobj.Query("select * from tg_user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	users := []telebot.User{}
	for rows.Next() {
		user := telebot.User{}
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username,
			&user.LanguageCode, &user.IsBot, &user.CanJoinGroups, &user.CanReadMessages,
			&user.SupportsInline)
		if err != nil {
			log.Fatalln(err)
		}
		users = append(users, user)
	}
	return users
}
