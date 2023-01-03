package main

import (
	"fmt"
	"github.com/balqisgautama/okami-midtrans/config"
	"github.com/balqisgautama/okami-midtrans/config/server"
	"github.com/balqisgautama/okami-midtrans/http/router"
	"github.com/balqisgautama/okami-midtrans/seeder"
	"github.com/balqisgautama/okami-midtrans/util"
	"os"
)

func main() {
	var arguments = "development"
	args := os.Args
	if len(args) > 1 {
		arguments = args[1]
	}

	config.GenerateConfiguration(arguments)
	server.SetServerConfig()
	seeder.DBMigrate()
	util.InitializeLogger()

	err := server.ServerConfig.DBConnection.Ping()
	if err != nil {
		fmt.Println("Connecting failed (PostgreSQL)", err)
	}

	defer func() {
		err := server.ServerConfig.DBConnection.Close()
		if err != nil {
			fmt.Println("Connecting failed (PostgreSQL)", err)
		}
	}()

	router.ApiController(config.ApplicationConfiguration.GetServerPort())
}
