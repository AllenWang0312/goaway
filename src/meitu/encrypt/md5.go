package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func MD5Encode(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func MD5Decode(str string) string {
	decode, err := hex.DecodeString(str)
	if nil == err {
		return string(decode)
	}
	return ""
}
func md5EncodeV2(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func md5EncodeV3(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}
