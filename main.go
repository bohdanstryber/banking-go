package main

import (
	"github.com/bohdanstryber/banking-go/app"
	"github.com/bohdanstryber/banking-go/logger"
)

func main() {
	logger.Info("Starting app...")
	app.Start()
}