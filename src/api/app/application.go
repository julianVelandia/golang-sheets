package app

import (
	"fmt"

	"github.com/julianVelandia/golang-sheets/src/api/app/config"
)

func StartApp() error {
	logEnvironment()
	router := NewRouter()
	return router.Run(config.GetConfig().Port())
}

func logEnvironment() {
	fmt.Println("GO_PORT: ", config.GetConfig().Port())
	fmt.Println("GO_SCOPE: ", config.GetConfig().Port())

	fmt.Println("Starting golang-sheets")
}
