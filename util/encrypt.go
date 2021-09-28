package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
	"strings"
)

// GetRandomSalt 生成盐值
func GetRandomSalt() string {
	return strings.ReplaceAll(uuid.NewV4().String(), "-", "")
}

func EncryptPassword(password, salt string) string {
	mac := hmac.New(sha1.New, []byte(password+salt))
	mac.Write([]byte(salt))
	s := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
