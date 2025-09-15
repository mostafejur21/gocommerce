package cmd

import (
	"fmt"
	"go_ecommerce/config"
)

func Serve() {
	fmt.Println("FROM SERVE FUNCTION")
	config.GetConfig()
}
