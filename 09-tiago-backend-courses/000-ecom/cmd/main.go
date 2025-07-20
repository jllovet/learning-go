package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jllovet/learning-go/09-tiago-backend-courses/000-ecom/cmd/api"
	"github.com/jllovet/learning-go/09-tiago-backend-courses/000-ecom/config"
	"github.com/jllovet/learning-go/09-tiago-backend-courses/000-ecom/db"
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
	defer db.Close()
	initStorage(db)
	server := api.NewAPIServer(fmt.Sprintf(":%s", config.Envs.Port), db)
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
