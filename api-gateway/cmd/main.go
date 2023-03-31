package main

import (
	"log"

	"github.com/wishrem/goligoli/api-gateway/router"
	"github.com/wishrem/goligoli/api-gateway/service"
	"github.com/wishrem/goligoli/logger"
)

func main() {
	logger.Setup(logger.Debug, log.Default())
	service.Setup()
	router.Setup()
}
