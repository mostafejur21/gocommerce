package cmd

import (
	"go_ecommerce/config"
	"go_ecommerce/rest"
	"go_ecommerce/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	middlewares := middlewares.NewMiddlewares(cnf)

	server := rest.NewServer(cnf)

	server.Start()
}
