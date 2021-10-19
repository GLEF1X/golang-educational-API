package main

import (
	"github.com/GLEF1X/golang-educational-API/api"
)

func main() {
	application := api.App{}
	application.Init()
	application.Run()
}
