package main

import (
	"fmt"

	"github.com/julianVelandia/GolangSheets/src/api/app"
)

func main() {
	if err := app.StartApp(); err != nil {
		fmt.Println("error starting server: ", err)
	}
}
