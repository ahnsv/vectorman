package main

import (
	"github.com/ahnsv/vectorman/router"
)

func main() {
	vectormanRouter := router.CreateRouter()
	vectormanRouter.Run()
}
