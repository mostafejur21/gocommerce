package cmd

import (
	"fmt"
	"go_ecommerce/config"
	"go_ecommerce/infra/db"
	"go_ecommerce/product"
	"go_ecommerce/repo"
	"go_ecommerce/rest"
	productHandler "go_ecommerce/rest/handlers/product"
	userHandler "go_ecommerce/rest/handlers/user"
	middleware "go_ecommerce/rest/middlewares"
	"go_ecommerce/user"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	middlewares := middleware.NewMiddlewares(cnf)

	// repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	// domains
	userSvc := user.NewService(userRepo)
	prdSvc := product.NewService(productRepo)

	productHandler := productHandler.NewHandler(middlewares, prdSvc)
	userHandler := userHandler.NewHandler(cnf, userSvc)

	server := rest.NewServer(cnf, productHandler, userHandler)

	server.Start()
}
