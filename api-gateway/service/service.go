package service

import "github.com/wishrem/goligoli/logger"

var Log = logger.Log

func Setup() {
	SetupUserClient()
	SetupVideoClient()
	SetupCommentClient()
}
