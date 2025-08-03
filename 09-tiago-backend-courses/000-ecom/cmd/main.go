package main

import (
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jllovet/learning-go/09-tiago-backend-courses/000-ecom/cmd/api"
	"github.com/jllovet/learning-go/09-tiago-backend-courses/000-ecom/config"
	"github.com/jllovet/learning-go/09-tiago-backend-courses/000-ecom/db"
)

func main() {
	cfg := config.InitializedConfig
	database, err := db.NewMySQLStorage(mysql.Config{
		User:                 cfg.DBUser,
		Passwd:               cfg.DBPassword,
		Addr:                 cfg.DBAddress,
		DBName:               cfg.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	db.InitStorage(database)
	server := api.NewAPIServer(fmt.Sprintf(":%s", cfg.Port), database)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
