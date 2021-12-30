package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/vpenando/piggy/pkg/routing"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()
	routing.HandleRoutes()
}
