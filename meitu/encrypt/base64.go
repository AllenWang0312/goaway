package encrypt

import (
	"crypto/rand"
	"fmt"
)

// 生成num*2位的字符串
func RandToken(num int) string {
	b := make([]byte, num)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
