package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/vpenando/piggy/routing"
)

var (
	database *gorm.DB
)

func init() {
	var err error
	database, err = gorm.Open(sqlite.Open(serverDatabase), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to init database: %s", err))
	}
	routing.InitFromConfig(currentLanguage, serverPort)
	routing.InitControllers(database)
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()
	routing.HandleRoutes()
}
