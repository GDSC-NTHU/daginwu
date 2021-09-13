package main

import (
	"kvlite/pkg/http"
	"kvlite/pkg/kv_agent"
	"log"
)

type App struct {
	kvAgent    kv_agent.KVAgent
	httpServer *http.HTTPServer
}

func (a *App) New() {

	a.kvAgent = &kv_agent.KVPebble{}
	a.httpServer = &http.HTTPServer{}
}

func (a *App) Init() {
	err := a.kvAgent.Init()
	if err != nil {
		log.Fatal("KVLITE init Pebble fail")
	}

	a.httpServer.Init(a.kvAgent)
}

func (a *App) Shutdown() {
	a.httpServer.Shutdown()
	log.Println("Shutdown HTTP Server")

	a.kvAgent.Shutdown()
	log.Println("Shutdown KV storage")
}

func (a *App) Run() {

	a.httpServer.Listen(":9999")
}
