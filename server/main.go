package main

import (
	"fmt"

	"github.com/ksharma67/EasyWay/server/app"
	"github.com/ksharma67/EasyWay/server/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.DBMigrate()
	fmt.Println("Working server on port:3000")
	app.Run(":3000")
}
