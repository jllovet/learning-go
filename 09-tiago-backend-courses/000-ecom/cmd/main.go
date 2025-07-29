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
	var Envs = config.InitConfig()
	database, err := db.NewMySQLStorage(mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPassword,
		Addr:                 Envs.DBAddress,
		DBName:               Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	db.InitStorage(database)
	server := api.NewAPIServer(fmt.Sprintf(":%s", Envs.Port), database)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
