package main

import (
	"fasthttp_restful/api"
)

func main() {
	application := api.App{}
	application.Init()
	application.Run()
}
