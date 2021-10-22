package main

import (
	"github.com/GLEF1X/golang-educational-API/api"
	"github.com/GLEF1X/golang-educational-API/core"
)

func main() {
	server := core.NewServer()
	api.SetupRoutes(server)
	server.Run()
}
