// Package ramdominfo 提供随机信息生成服务
package ramdominfo

import (
	"crypto/md5"
	cryptorand "crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"strings"
)

// RamdomNumber 返回[min, max]随机数字
func RamdomNumber(min, max uint64) (v uint64) {
	if min > max {
		min, max = max, min
	}
	binary.Read(cryptorand.Reader, binary.BigEndian, &v)
	return v%(max-min) + min
}

// RamdomBytes 随机字节数组
func RamdomBytes(n int) []byte {
	b := make([]byte, n)
	cryptorand.Read(b)
	return b
}

// RamdomMD5String 随机md5字符串
func RamdomMD5String(n int) string {
	m := md5.New()
	m.Write(RamdomBytes(n))
	return hex.EncodeToString(m.Sum(nil))
}

// RamdomMD5UpperString 随机md5字符串, 大写
func RamdomMD5UpperString(n int) string {
	return strings.ToUpper(RamdomMD5String(n))
}
