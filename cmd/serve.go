package cmd

import (
	"fmt"
	"go_ecommerce/config"
	"go_ecommerce/infra/db"
	"go_ecommerce/repo"
	"go_ecommerce/rest"
	"go_ecommerce/rest/middlewares"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	middlewares := middlewares.NewMiddlewares(cnf)

	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	server := rest.NewServer(cnf)

	server.Start()
}
