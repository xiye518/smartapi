package tools

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func GetGuid() string {
	a := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, a); err != nil {
		return ""
	}
	s := GetMd5String(base64.URLEncoding.EncodeToString(a))
	b := []byte(s)
	//r := fmt.Sprintf("%s-%s-%s-%s-%s", b[0:8], b[8:12], b[12:16], b[16:20], b[20:])
	return fmt.Sprintf("%s-%s-%s-%s-%s", b[0:8], b[8:12], b[12:16], b[16:20], b[20:])
	//return GetMd5String(base64.URLEncoding.EncodeToString(a))
}
