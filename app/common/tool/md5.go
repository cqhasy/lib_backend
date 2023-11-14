package tool

import (
	MD5 "crypto/md5"
	"encoding/hex"
)

func EncryptedPasswordMD5(password string) string {
	md5 := MD5.New()
	md5.Write([]byte(password))
	HashPassword := hex.EncodeToString(md5.Sum(nil))
	return HashPassword
}
