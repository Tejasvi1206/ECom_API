package main

import (
	"database/sql"
	"log"

	"github.com/Tejasvi1206/ECom_API/cmd/api"
	"github.com/Tejasvi1206/ECom_API/config"
	"github.com/Tejasvi1206/ECom_API/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db) //Confirms DB is reachable at startup and Avoids runtime crashes later

	server := api.NewAPIServer(":8080", db) // here we are Passing DB into the server NOT creating DB inside the server Making the server testable This is called Dependency Injection This is advanced backend design, not beginner.
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
