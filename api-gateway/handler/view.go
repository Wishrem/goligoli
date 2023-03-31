package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/wishrem/goligoli/pkg/conf"
)

type File struct {
	FileType string `uri:"type" binding:"required"`
	FileName string `uri:"name" binding:"required"`
}

func View(c *gin.Context) {
	file := new(File)
	if err := c.ShouldBindUri(&file); err != nil {
		SendBadRequest(c)
		return
	}

	dir, ok := getDir(file.FileType)
	if !ok {
		c.Redirect(http.StatusNotFound, "/goligoli/404")
		return
	}

	filename := dir + file.FileName
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusFound, "/goligoli/404")
		return
	}
	defer f.Close()

	c.File(filename)
}

func getDir(fileType string) (string, bool) {
	switch fileType {
	case "photo":
		return conf.App.Res.PhotoDir, true
	case "video":
		return conf.App.Res.VideoDir, true
	default:
		return "", false
	}
}
