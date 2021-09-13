package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type AppCtl struct {
	app *App
}

func (ctl *AppCtl) Init() {

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Application ...")
	ctl.app.Shutdown()

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	time.Sleep(5 * time.Second)
	defer cancel()
	log.Println("Exiting Application ...")

}
