package main

import (
	application "CalcServer/application"
)

func main() {
	app := application.NewApp()
	app.StartServer()
}
