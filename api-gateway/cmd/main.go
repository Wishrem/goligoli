package main

import (
	"github.com/wishrem/goligoli/api-gateway/internal/service"
	"github.com/wishrem/goligoli/api-gateway/router"
)

func main() {
	service.Setup()
	router.Setup()
}
