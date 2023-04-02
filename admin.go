package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/wishrem/goligoli/pkg/util/snowflake"
	"github.com/wishrem/goligoli/user/model"
	"github.com/yitter/idgenerator-go/idgen"
)

func main() {
	model.Init()
	snowflake.Init(0)
	if err := CreateAdmin(); err != nil {
		panic(err)
	}
}

func CreateAdmin() error {
	u := new(model.User)
	u.ID = idgen.NextId()
	u.Roles = []*model.Role{
		{
			Type: "admin",
		},
		{
			Type: "user",
		},
	}
	u.Description = "admin"
	u.Name = "admin"
	u.Email = "admin@goligoli.com"
	u.PhotoUrl = "127.0.0.1:8080/goligoli/view/photo/default.jpg"
	s, err := getSHA256String("admin")
	if err != nil {
		return err
	}
	u.Password = s

	return u.Create(context.Background())
}

func getSHA256String(s string) (string, error) {
	msg := []byte(s)
	hash := sha256.New()
	n, err := hash.Write(msg)
	if err != nil {
		return "", err
	}

	if n != len(msg) {
		return "", errors.New("encoder password error")
	}
	bytes := hash.Sum(nil)
	return hex.EncodeToString(bytes), nil
}
