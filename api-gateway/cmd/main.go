package main

import (
	"github.com/wishrem/goligoli/api-gateway/router"
	"github.com/wishrem/goligoli/api-gateway/service"
)

func main() {
	service.Setup()
	router.Setup()
}
