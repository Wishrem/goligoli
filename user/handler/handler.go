package handler

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/wishrem/goligoli/user/model"
	"github.com/wishrem/goligoli/user/proto/pb"
)

func parseUser(user *model.User) *pb.User {
	u := new(pb.User)
	u.Ban = parseBan(&user.Ban)
	u.Roles = parseRole(user.Roles...)

	u.Id = user.ID
	u.Name = user.Name
	u.Email = user.Email
	u.PhotoUrl = user.PhotoUrl
	u.Description = user.Description
	return u
}

func parseBan(ban *model.Ban) *pb.Ban {
	b := new(pb.Ban)
	b.BanAt = ban.BanAt.Unix()
	b.Reason = ban.Reason
	b.UnbanAt = ban.UnbanAt.Unix()
	return b
}

func parseRole(role ...model.Role) []string {
	r := make([]string, 0, len(role))
	for _, s := range role {
		r = append(r, s.Type)
	}
	return r
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
