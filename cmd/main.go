package main

import (
	"github.com/Moriartii/translator/pkg/starter"
)

func start() {
	app := starter.NewApplication()
	app.Run()
}

func main() {
	start()
}
