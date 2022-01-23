package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/vpenando/piggy/pkg/config"
	"github.com/vpenando/piggy/pkg/piggy"
	"github.com/vpenando/piggy/pkg/routing"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()
	log.Println("Current version:", piggy.Version)
	config.ReadConfig()
	routing.HandleRoutes()
}
