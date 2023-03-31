package util

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
)

func GenerateFileName(bytes []byte) (string, error) {
	h := md5.New()
	n, err := h.Write(bytes)
	if err != nil {
		return "", err
	}
	if n != len(bytes) {
		return "", errors.New("encoder file name error")
	}
	b := h.Sum(nil)
	return hex.EncodeToString(b), nil
}
